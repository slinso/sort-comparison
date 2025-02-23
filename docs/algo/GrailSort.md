# Grail Sort

Grail Sort is a stable in-place merge sort variant that uses O(√n) additional memory. It was developed by Andrey Astrelin as a solution to create a highly efficient, stable sorting algorithm with minimal additional memory requirements. For more details on the algorithm, see the [Grail Sort implementation](https://github.com/BonzaiThePenguin/WikiSort/blob/master/grailsort.h).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/GrailSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Grail Sort maintains O(n log n) complexity in all cases while being stable and using only O(√n) additional memory, making it particularly useful when memory is constrained but stability is required.