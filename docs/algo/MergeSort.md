# Merge Sort

Merge Sort is a stable, divide-and-conquer sorting algorithm that recursively divides the input array into smaller subarrays, sorts them, and then merges them back together. For more details on the algorithm and its theory, see the [Merge Sort Wikipedia article](https://en.wikipedia.org/wiki/Merge_sort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/MergeSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Merge Sort consistently performs in O(n log n) time for all cases, making it reliable for any input. While it requires O(n) additional space, its stability and predictable performance make it valuable for sorting linked lists and in external sorting scenarios where data doesn't fit entirely in memory.