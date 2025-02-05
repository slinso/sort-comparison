# Bucket Sort

Bucket Sort is a distribution-based sorting algorithm that works by distributing elements into a number of buckets, then sorting these buckets individually. Each bucket holds a range of values, and the final sorted array is produced by concatenating all buckets in order. For more details on the algorithm and its theory, see the [Bucket Sort Wikipedia article](https://en.wikipedia.org/wiki/Bucket_sort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                  |
| ------------------ | ---------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/BucketSort_cat_d_series_s_10$_bars.svg" width="600">    |
| 100                | <img src="../../images/perf/algo/BucketSort_cat_d_series_s_100$_bars.svg" width="600">   |
| 1,000              | <img src="../../images/perf/algo/BucketSort_cat_d_series_s_1000$_bars.svg" width="600">  |
| 10,000             | <img src="../../images/perf/algo/BucketSort_cat_d_series_s_10000$_bars.svg" width="600"> |

Note: While Bucket Sort can achieve O(n+k) complexity in the best and average cases, its worst-case performance of O(n²) and memory requirements limit its practical use to smaller datasets.