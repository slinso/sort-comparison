# Bubble Sort

Bubble Sort is a simple comparison-based sorting algorithm. It repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order. The pass through the list is repeated until no swaps are needed. For more details on the algorithm and its theory, see the [Bubble Sort Wikipedia article](https://en.wikipedia.org/wiki/Bubble_sort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                  |
| ------------------ | ---------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/BubbleSort_cat_d_series_s_10$_bars.svg" width="600">    |
| 100                | <img src="../../images/perf/algo/BubbleSort_cat_d_series_s_100$_bars.svg" width="600">   |
| 1,000              | <img src="../../images/perf/algo/BubbleSort_cat_d_series_s_1000$_bars.svg" width="600">  |
| 10,000             | <img src="../../images/perf/algo/BubbleSort_cat_d_series_s_10000$_bars.svg" width="600"> |

Note: Being an O(n²) algorithm, Bubble Sort's performance becomes impractical for larger datasets, which is why benchmarks are limited to 10,000 elements.