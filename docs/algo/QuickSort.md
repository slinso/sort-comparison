# Quick Sort

Quick Sort is a highly efficient, comparison-based sorting algorithm that uses a divide-and-conquer strategy. It works by selecting a 'pivot' element and partitioning the array around it, such that smaller elements are moved before the pivot and larger elements after it. For more details on the algorithm and its theory, see the [Quick Sort Wikipedia article](https://en.wikipedia.org/wiki/Quicksort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/QuickSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Quick Sort achieves O(n log n) complexity in best and average cases but can degrade to O(n²) in the worst case (when the array is already sorted or reverse sorted). It requires O(log n) additional space for recursion. Despite its worst-case behavior, Quick Sort is often the fastest sorting algorithm in practice due to its excellent cache performance and ability to sort in-place.