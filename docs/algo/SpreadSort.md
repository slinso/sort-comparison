# Spread Sort

Spread Sort is a hybrid sorting algorithm that combines the distribution-based approach with comparison sorting. It spreads out elements based on their values and uses this information to sort efficiently. For more details on the algorithm, see the [Spread Sort implementation](https://github.com/orlp/pdqsort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                      |
| ------------------ | -------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/SpreadSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Spread Sort achieves O(n) complexity in the best case and O(n log² n) in average and worst cases. It requires O(n) additional memory. The algorithm is particularly effective when dealing with data that has a relatively uniform distribution.