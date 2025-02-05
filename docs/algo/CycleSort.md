# Cycle Sort

Cycle Sort is an in-place, unstable sorting algorithm that is theoretically optimal in terms of the number of memory writes. It works by cycling through each element and placing it in its correct position, creating a series of cyclic rotations. For more details on the algorithm and its theory, see the [Cycle Sort Wikipedia article](https://en.wikipedia.org/wiki/Cycle_sort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                 |
| ------------------ | --------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/CycleSort_cat_d_series_s_10$_bars.svg" width="600">    |
| 100                | <img src="../../images/perf/algo/CycleSort_cat_d_series_s_100$_bars.svg" width="600">   |
| 1,000              | <img src="../../images/perf/algo/CycleSort_cat_d_series_s_1000$_bars.svg" width="600">  |
| 10,000             | <img src="../../images/perf/algo/CycleSort_cat_d_series_s_10000$_bars.svg" width="600"> |

Note: While Cycle Sort minimizes the number of memory writes, its O(n²) time complexity makes it primarily useful for situations where write operations are significantly more expensive than read operations or where minimal memory writes are crucial.