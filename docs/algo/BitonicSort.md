# Bitonic Sort

Bitonic Sort is a parallel algorithm for sorting that is particularly suited for hardware implementation. It works by repeatedly merging bitonic sequences until the entire sequence is sorted. A bitonic sequence is one that first monotonically increases, then monotonically decreases. For more details on the algorithm and its theory, see the [Bitonic Sort Wikipedia article](https://en.wikipedia.org/wiki/Bitonic_sorter).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/BitonicSort_cat_d_series_s_10$_bars.svg" width="600">      |
| 100                | <img src="../../images/perf/algo/BitonicSort_cat_d_series_s_100$_bars.svg" width="600">     |
| 1,000              | <img src="../../images/perf/algo/BitonicSort_cat_d_series_s_1000$_bars.svg" width="600">    |
| 10,000             | <img src="../../images/perf/algo/BitonicSort_cat_d_series_s_10000$_bars.svg" width="600">   |
| 100,000            | <img src="../../images/perf/algo/BitonicSort_cat_d_series_s_100000$_bars.svg" width="600">  |
| 1,000,000          | <img src="../../images/perf/algo/BitonicSort_cat_d_series_s_1000000$_bars.svg" width="600"> |