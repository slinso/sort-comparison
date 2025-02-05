# Weave Merge Sort

Weave Merge Sort is a variation of merge sort that uses a weaving pattern to combine sorted sequences. It focuses on efficient memory usage and cache performance by interleaving elements during the merge phase. For more details on merge sort variations, see the [Merge Sort Wikipedia article](https://en.wikipedia.org/wiki/Merge_sort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                          |
| ------------------ | ------------------------------------------------------------------------------------------------ |
| 10                 | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/WeaveMergeSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Weave Merge Sort achieves O(n log n) complexity in all cases and requires O(n) additional memory. The algorithm maintains stability and can provide better cache performance than traditional merge sort implementations due to its weaving merge pattern.