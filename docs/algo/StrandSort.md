# Strand Sort

Strand Sort is a sorting algorithm that works by repeatedly pulling sorted subsequences (strands) out of the input list and merging them with a result array. It identifies already sorted sequences within the input and uses them to build the final sorted output. For more details on the algorithm and its theory, see the [Strand Sort description](https://en.wikipedia.org/wiki/Strand_sort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                    |
| ------------------ | ------------------------------------------------------------------------------------------ |
| 10                 | <img src="../../images/perf/algo/StrandSort_cat_d_series_s_10$_bars.svg" width="600">      |
| 100                | <img src="../../images/perf/algo/StrandSort_cat_d_series_s_100$_bars.svg" width="600">     |
| 1,000              | <img src="../../images/perf/algo/StrandSort_cat_d_series_s_1000$_bars.svg" width="600">    |
| 10,000             | <img src="../../images/perf/algo/StrandSort_cat_d_series_s_10000$_bars.svg" width="600">   |
| 100,000            | <img src="../../images/perf/algo/StrandSort_cat_d_series_s_100000$_bars.svg" width="600">  |
| 1,000,000          | <img src="../../images/perf/algo/StrandSort_cat_d_series_s_1000000$_bars.svg" width="600"> |

Note: Strand Sort achieves O(n) complexity in the best case (when the input contains few reverse ordered subsequences) but degrades to O(n²) in average and worst cases. It requires O(n) additional memory and maintains stability. The algorithm is particularly efficient when the input contains many sorted subsequences.