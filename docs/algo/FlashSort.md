# Flash Sort

Flash Sort is a distribution sorting algorithm that uses a permutation to produce a sorted sequence. It is particularly efficient when dealing with uniformly distributed data. For more details on the algorithm and its theory, see the [Flash Sort paper](https://dl.acm.org/doi/10.1145/178243.178259).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/FlashSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Flash Sort achieves O(n) complexity in the best and average cases but can degrade to O(n²) in the worst case. It requires O(n) additional memory.