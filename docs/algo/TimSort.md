# Tim Sort

Tim Sort is a hybrid stable sorting algorithm, combining merge sort and insertion sort, designed to perform well on many kinds of real-world data. It was implemented by Tim Peters in 2002 for Python's sorting method. For more details on the algorithm and its theory, see the [Tim Sort Wikipedia article](https://en.wikipedia.org/wiki/Timsort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                   |
| ------------------ | ----------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/TimSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/TimSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/TimSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/TimSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/TimSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/TimSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/TimSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/TimSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Tim Sort achieves O(n) complexity in the best case (already sorted) and O(n log n) in average and worst cases. It requires O(n) additional memory and is stable. The algorithm is particularly efficient on real-world data that often contains partially ordered subsequences.