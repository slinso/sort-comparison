package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/perf/benchfmt"
)

type Result struct {
	Name  string
	Iters int
	Value []benchfmt.Value
}

func main() {
	flag.Parse()

	stat := []Result{}
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

			stat = append(stat, res)
		}
	}
	if err := files.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, s := range stat {
		fmt.Println(s.Name, s.Iters, s.Value)
	}
}
