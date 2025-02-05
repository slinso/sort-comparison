# Intro Sort

Intro Sort or Introspective Sort is a hybrid sorting algorithm that provides both fast average performance and optimal worst-case performance. It begins with quicksort and switches to heapsort when the recursion depth exceeds a certain level, and to insertion sort for small subarrays. For more details on the algorithm and its theory, see the [Introsort Wikipedia article](https://en.wikipedia.org/wiki/Introsort).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_10$_bars.svg" width="600">        |
| 100                | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_100$_bars.svg" width="600">       |
| 1,000              | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_1000$_bars.svg" width="600">      |
| 10,000             | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_10000$_bars.svg" width="600">     |
| 100,000            | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_100000$_bars.svg" width="600">    |
| 1,000,000          | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_1000000$_bars.svg" width="600">   |
| 10,000,000         | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_10000000$_bars.svg" width="600">  |
| 100,000,000        | <img src="../../images/perf/algo/IntroSort_cat_d_series_s_100000000$_bars.svg" width="600"> |

Note: Intro Sort guarantees O(n log n) worst-case performance while keeping quicksort's average-case efficiency. It requires O(log n) additional space for recursion. This algorithm is used in many standard library implementations, including C++'s std::sort.