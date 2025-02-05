# Block Sort

Block Sort is a sorting algorithm that combines elements of insertion sort with block operations for improved efficiency. It works by dividing the input into blocks, sorting them individually, and then merging them in a way that minimizes memory operations. For more details on the algorithm and its theory, see the [Block Sort description](https://en.wikipedia.org/wiki/Block_sort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                   |
| ------------------ | ----------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/BlockSort_cat_d_series_s_10$_bars.svg" width="600">      |
| 100                | <img src="../../images/perf/algo/BlockSort_cat_d_series_s_100$_bars.svg" width="600">     |
| 1,000              | <img src="../../images/perf/algo/BlockSort_cat_d_series_s_1000$_bars.svg" width="600">    |
| 10,000             | <img src="../../images/perf/algo/BlockSort_cat_d_series_s_10000$_bars.svg" width="600">   |
| 100,000            | <img src="../../images/perf/algo/BlockSort_cat_d_series_s_100000$_bars.svg" width="600">  |
| 1,000,000          | <img src="../../images/perf/algo/BlockSort_cat_d_series_s_1000000$_bars.svg" width="600"> |