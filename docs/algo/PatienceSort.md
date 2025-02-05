# Patience Sort

Patience Sort is a sorting algorithm inspired by the card game patience/solitaire. It works by creating piles of cards (elements) where each pile is sorted, similar to how players create piles in solitaire. For more details on the algorithm and its theory, see the [Patience Sort Wikipedia article](https://en.wikipedia.org/wiki/Patience_sorting).

## Benchmark Results

| Number of Elements | Benchmark Visualization                                                                      |
| ------------------ | -------------------------------------------------------------------------------------------- |
| 10                 | <img src="../../images/perf/algo/PatienceSort_cat_d_series_s_10$_bars.svg" width="600">      |
| 100                | <img src="../../images/perf/algo/PatienceSort_cat_d_series_s_100$_bars.svg" width="600">     |
| 1,000              | <img src="../../images/perf/algo/PatienceSort_cat_d_series_s_1000$_bars.svg" width="600">    |
| 10,000             | <img src="../../images/perf/algo/PatienceSort_cat_d_series_s_10000$_bars.svg" width="600">   |
| 100,000            | <img src="../../images/perf/algo/PatienceSort_cat_d_series_s_100000$_bars.svg" width="600">  |
| 1,000,000          | <img src="../../images/perf/algo/PatienceSort_cat_d_series_s_1000000$_bars.svg" width="600"> |

Note: Patience Sort achieves O(n log n) complexity in all cases and is stable. While it requires O(n) additional space, it has the unique property of being able to determine the longest increasing subsequence in the input as a side effect of the sorting process.