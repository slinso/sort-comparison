# Sort Comparison

A comprehensive collection of sorting algorithm implementations in Go for
educational purposes.

## Algorithms Implemented

### Comparison Based Sorts

- BubbleSort - O(n²), stable, in-place
- CocktailShakerSort - O(n²), stable, in-place
- CombSort - O(n log n), not stable, in-place
- ExchangeSort - O(n²), not stable, in-place
- GnomeSort - O(n²), stable, in-place
- HeapSort - O(n log n), not stable, in-place
- InsertionSort - O(n²), stable, in-place
- IntroSort - O(n log n), not stable, in-place
- MergeSort - O(n log n), stable, O(n) space
- QuickSort - O(n log n), not stable, in-place
- SelectionSort - O(n²), not stable, in-place
- ShellSort - O(n log² n), not stable, in-place
- TimSort - O(n log n), stable, O(n) space

### Distribution Based Sorts

- BlockSort - O(n log n), stable, O(n) space
- BucketSort - O(n + k), stable, O(n + k) space
- CountingSort - O(n + k), stable, O(k) space
- FlashSort - O(n), not stable, O(n) space
- PigeonholeSort - O(n + range), stable, O(range) space
- RadixSort - O(d \* (n + k)), stable, O(n + k) space
- SpreadSort - O(n log log n), not stable, O(n) space

### Specialized Sorts

- BurstSort - O(n log n), stable, O(n) space
- CycleSort - O(n²), not stable, optimal memory writes
- LibrarySort - O(n log n), stable, O(n) space
- PatienceSort - O(n log n), stable, O(n) space
- PostmanSort - O(n + r), stable, O(n) space
- SimplePancakeSort - O(n²), not stable, in-place
- SmoothSort - O(n log n), not stable, in-place
- StrandSort - O(n²), stable, O(n) space
- TournamentSort - O(n log n), not stable, O(n) space
- TreeSort - O(n log n), stable, O(n) space

## Benchmarking

```go
go test -bench=.
```

Compare specific algorithms:

```go
go test -bench "//Quick/10"
```

Benchmarks test each algorithm against different:

- Input sizes (10 to 1B elements)
- Data distributions (random, sorted, reversed, etc.)
- Operation types (in-place vs return new array)

Usage License This project is licensed under the MIT License.
