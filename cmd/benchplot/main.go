package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/aclements/go-moremath/stats"
	"github.com/samber/lo"
	"github.com/vicanso/go-charts/v2"
	"golang.org/x/perf/benchfmt"
)

type Result struct {
	Name  string
	Iters int
	Value []benchfmt.Value
}

func main() {
	flagTable := flag.String("table", "", "[a|d|s] split data into multiple charts")
	flagCategories := flag.String("categories", "", "[a|d|s]")
	flagSeries := flag.String("series", "", "[a|d|s]")
	flagAvg := flag.Bool("avg", false, "average per category")
	flagFilter := flag.String("filter", "", "filter benchmarks (regexp pattern)")

	flag.Parse()

	// check flags
	if *flagCategories == "" {
		panic("categories not set")
	}

	if *flagSeries == "" {
		panic("series not set")
	}

	stat := map[string][]Result{}
	files := benchfmt.Files{Paths: flag.Args(), AllowStdin: true, AllowLabels: true}
	for files.Scan() {
		switch rec := files.Result(); rec := rec.(type) {
		case *benchfmt.SyntaxError:
			// Non-fatal result parse error. Warn
			// but keep going.
			fmt.Fprintln(os.Stderr, rec)
		case *benchfmt.Result:
			res := Result{
				Name:  string(rec.Name),
				Iters: rec.Iters,
			}
			res.Value = make([]benchfmt.Value, len(rec.Values))
			copy(res.Value, rec.Values)

			stat[res.Name] = append(stat[res.Name], res)
		}
	}
	if err := files.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fromFlag := func(b BenchmarkMeta, flagValue string) string {
		switch flagValue {
		case "a":
			return b.Algorithm
		case "d":
			return b.Distribution
		case "s":
			return b.Size
		default:
			return ""
		}
	}

	benchmarks := []BenchmarkMeta{}

	for key, value := range stat {
		meta, err := parseBenchmarkName(key)
		if err != nil {
			panic(err)
		}

		values := lo.Map(value, func(val Result, _ int) float64 {
			return lo.Reduce(val.Value, func(acc float64, v benchfmt.Value, _ int) float64 {
				if v.Unit == "sec/op" {
					return v.Value
				}
				return acc
			}, 0)
		})
		mean, lo, hi := stats.MeanCI(values, 0.95)
		meta.Mean = mean
		meta.DiffMean = meanDiff(mean, lo, hi)

		meta.Table = fromFlag(meta, *flagTable)
		meta.Category = fromFlag(meta, *flagCategories)
		meta.Series = fromFlag(meta, *flagSeries)

		if *flagFilter != "" {
			re, err := regexp.Compile(*flagFilter)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid regexp pattern: %v\n", err)
				os.Exit(1)
			}
			if re.MatchString(key) {
				benchmarks = append(benchmarks, meta)
			}
		} else {
			benchmarks = append(benchmarks, meta)
		}
	}

	tables := lo.UniqMap(benchmarks, func(b BenchmarkMeta, _ int) string {
		return b.Table
	})
	slices.Sort(tables)

	categories := lo.UniqMap(benchmarks, func(b BenchmarkMeta, _ int) string {
		return b.Category
	})
	slices.Sort(categories)

	series := lo.UniqMap(benchmarks, func(b BenchmarkMeta, _ int) string {
		return b.Series
	})
	slices.Sort(series)

	plots := make(map[string]BenchmarkPlot)

	for _, b := range benchmarks {
		// create if not exists
		if _, ok := plots[b.Table]; !ok {
			plot := BenchmarkPlot{
				Title:  b.Table,
				Values: make(map[string]map[string][]float64),
			}
			plots[b.Table] = plot
			// init categories
			for _, c := range categories {
				plots[b.Table].Values[c] = make(map[string][]float64)
				// init series
				for _, s := range series {
					plots[b.Table].Values[c][s] = make([]float64, 0)
				}
			}
		}

		// add data
		plots[b.Table].Values[b.Category][b.Series] = append(plots[b.Table].Values[b.Category][b.Series], b.Mean*1e9)

	}

	for _, t := range tables {
		values := make([][]float64, 0)

		if *flagAvg {
			row := make([]float64, 0)
			for _, c := range categories {
				row = append(row, lo.Reduce(series, func(acc float64, s string, _ int) float64 {
					return acc + plots[t].Values[c][s][0]
				}, 0)/float64(len(series)))
			}
			values = append(values, row)
		} else {
			for _, s := range series {
				row := make([]float64, 0)
				for _, c := range categories {
					row = append(row, plots[t].Values[c][s]...)
				}
				values = append(values, row)
			}
		}

		output := "svg"

		p, err := charts.BarRender(
			values,
			charts.PaddingOptionFunc(charts.Box{Top: 10, Right: 50, Bottom: 100, Left: 10}),
			charts.WidthOptionFunc(2400),
			charts.HeightOptionFunc(800),
			charts.LegendOptionFunc(charts.LegendOption{
				Data: lo.Ternary(*flagAvg, []string{"Average"}, series),
				Padding: charts.Box{
					Top:  30,
					Left: 50,
				},
			}),
			charts.TypeOptionFunc(output),
			charts.TitleOptionFunc(charts.TitleOption{Text: t, FontSize: 20}),
			charts.XAxisDataOptionFunc(categories),
			charts.YAxisOptionFunc(
				charts.YAxisOption{
					Formatter: "{value} ns",
				},
			),
			charts.MarkLineOptionFunc(0, charts.SeriesMarkDataTypeAverage),
			charts.MarkPointOptionFunc(0, charts.SeriesMarkDataTypeMax, charts.SeriesMarkDataTypeMin),
			// custom option func
			func(opt *charts.ChartOption) {
				opt.XAxis = charts.XAxisOption{
					Data:         categories,
					TextRotation: 4.8,
					LabelOffset: charts.Box{
						Top: -5,
					},
					FontSize: 10,
				}
				opt.SeriesList[0].MarkPoint = charts.NewMarkPoint(
					charts.SeriesMarkDataTypeMax,
					charts.SeriesMarkDataTypeMin,
				)
				opt.SeriesList[0].MarkLine = charts.NewMarkLine(
					charts.SeriesMarkDataTypeAverage,
				)
			},
		)
		if err != nil {
			panic(err)
		}

		buf, err := p.Bytes()
		if err != nil {
			panic(err)
		}
		err = writeFile(buf, t, *flagCategories, *flagSeries, *flagFilter, *flagAvg, output)
		if err != nil {
			panic(err)
		}
	}
}

func writeFile(buf []byte, tablename string, categorie string, series string, filter string, avg bool, output string) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0o700)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s_cat_%s_series_%s", tablename, categorie, series)
	if avg {
		filename += "_avg"
	}
	if filter != "" {
		filename += fmt.Sprintf("_%s", filter)
	}
	filename += fmt.Sprintf("_bars.%s", output)

	file := filepath.Join(tmpPath, filename)
	err = os.WriteFile(file, buf, 0o600)
	if err != nil {
		return err
	}
	return nil
}

type BenchmarkPlot struct {
	Title  string
	Values map[string]map[string][]float64
}

// 1. Algo, GroupBy: distribution, X: datapoints, Y: time
// 2. Distributiong, GroupBy: datapoints, X: algo, Y: time
// 3. Datapoints, GroupBy: algo, X: distribution, Y: time
// 4. Average per Algo, GroupBy: "", X: algo, Y: time

type BenchmarkMeta struct {
	Distribution string
	Algorithm    string
	Size         string
	Mean         float64
	DiffMean     float64

	Table    string
	Category string
	Series   string
}

func parseBenchmarkName(name string) (BenchmarkMeta, error) {
	parts := strings.Split(name, "/")
	meta := BenchmarkMeta{}

	for _, part := range parts[1:] { // Skip first part ("Sort")
		kv := strings.Split(part, "=")
		if len(kv) != 2 {
			continue
		}

		switch kv[0] {
		case "dist":
			meta.Distribution = kv[1]
		case "algo":
			meta.Algorithm = kv[1]
		case "size":
			meta.Size = kv[1]
		}
	}

	return meta, nil
}

type Scaler func(float64) string

func timeScaler(ns float64) Scaler {
	var format string
	var scale float64
	switch x := ns; {
	case x >= 99.5:
		format, scale = "%.0fs", 1
	case x >= 9.95:
		format, scale = "%.1fs", 1
	case x >= 0.995:
		format, scale = "%.2fs", 1
	case x >= 0.0995:
		format, scale = "%.0fms", 1000
	case x >= 0.00995:
		format, scale = "%.1fms", 1000
	case x >= 0.000995:
		format, scale = "%.2fms", 1000
	case x >= 0.0000995:
		format, scale = "%.0fµs", 1000*1000
	case x >= 0.00000995:
		format, scale = "%.1fµs", 1000*1000
	case x >= 0.000000995:
		format, scale = "%.2fµs", 1000*1000
	case x >= 0.0000000995:
		format, scale = "%.0fns", 1000*1000*1000
	case x >= 0.00000000995:
		format, scale = "%.1fns", 1000*1000*1000
	default:
		format, scale = "%.2fns", 1000*1000*1000
	}
	return func(ns float64) string {
		return fmt.Sprintf(format, ns*scale)
	}
}

// FormatDiff computes and formats the percent variation of max and min compared to mean.
// If b.Mean or b.Max is zero, FormatDiff returns an empty string.
func meanDiff(mean float64, lo float64, hi float64) float64 {
	if mean == 0 || hi == 0 {
		return 0
	}
	diff := 1 - lo/mean
	if d := hi/mean - 1; d > diff {
		diff = d
	}
	return diff
}
