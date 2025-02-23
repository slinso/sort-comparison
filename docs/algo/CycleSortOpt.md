# CycleSortOpt

CycleSortOpt is a variant of cycle sort that optimizes the number of writes during sorting. It is designed to minimize the data movements required, making it efficient in scenarios where write operations are costly. For more details on the algorithm and its theory, see the [Sorting algorithm Wikipedia article](https://en.wikipedia.org/wiki/Sorting_algorithm).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                         |
| ------------------ | ----------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_10$_bars.svg" width="600">         |
| 100                | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_100$_bars.svg" width="600">        |
| 1,000              | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_1000$_bars.svg" width="600">       |
| 10,000             | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_10000$_bars.svg" width="600">      |
| 100,000            | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_100000$_bars.svg" width="600">     |
| 1,000,000          | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_1000000$_bars.svg" width="600">    |
| 10,000,000         | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_10000000$_bars.svg" width="600">   |
| 100,000,000        | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_100000000$_bars.svg" width="600">  |
| 1,000,000,000      | <img src="../../images/perf/algo/CycleSortOpt_cat_d_series_s_1000000000$_bars.svg" width="600"> |