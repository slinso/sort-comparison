# Sort Comparison

A comprehensive collection of sorting algorithm implementations in Go for
educational purposes.

# ideas

- one svg bar chart for each algo
- one svg bar chart for each distribution
- svg generation might be simpler to incorporate into a local test setup
- a single gigantic echarts with distribution as series + on algo has to be
  selectable as well

# what benchmarks work

- StdSort up to 1B
- AdaptiveSort up to 100M
- AmericanFlagSort up to 100M
- BeadSort up to 10K - ohne Random
- BitonicSort up to 10M
- BlockSort up to 10M
- BubbleSort up to 100K
- BucketSort up to 100K
- BurstSort up to 100M
- CascadeSort up to 10M
- CocktailshakerSort up to 10K
- CombSort up to 10M
- CountingSort up to 100M

## Sorting Algorithms

### Comparison Based

| Name              | Best       | Average     | Worst       | Memory   | Stable | In-place | Method               | Notes                                           |
| ----------------- | ---------- | ----------- | ----------- | -------- | ------ | -------- | -------------------- | ----------------------------------------------- |
| **BubbleSort**    | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      | Comparison           | Simple implementation, good for teaching        |
| **CocktailSort**  | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      | Bidirectional bubble | Slightly better than bubble sort                |
| **CombSort**      | O(n log n) | O(n log n)  | O(n²)       | O(1)     | No     | Yes      | Gap sequence         | Improvement over bubble sort                    |
| **ExchangeSort**  | O(n²)      | O(n²)       | O(n²)       | O(1)     | No     | Yes      | Comparison           | Similar to bubble sort                          |
| **GnomeSort**     | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      | Insertion variant    | Simple but inefficient                          |
| **HeapSort**      | O(n log n) | O(n log n)  | O(n log n)  | O(1)     | No     | Yes      | Selection heap       | Optimal complexity but cache-unfriendly         |
| **InsertionSort** | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      | Insertion            | Excellent for small arrays                      |
| **IntroSort**     | O(n log n) | O(n log n)  | O(n log n)  | O(log n) | No     | Yes      | Hybrid quicksort     | Combines quicksort, heapsort and insertion sort |
| **MergeSort**     | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       | Divide & conquer     | Stable but requires extra space                 |
| **QuickSort**     | O(n log n) | O(n log n)  | O(n²)       | O(log n) | No     | Yes      | Partitioning         | Fast in practice, widely used                   |
| **SelectionSort** | O(n²)      | O(n²)       | O(n²)       | O(1)     | No     | Yes      | Selection            | Minimal memory writes                           |
| **ShellSort**     | O(n log n) | O(n log² n) | O(n log² n) | O(1)     | No     | Yes      | Gap insertion        | Good for medium-sized arrays                    |
| **TimSort**       | O(n)       | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       | Hybrid merge         | Python's standard sort                          |
| **OddEvenSort**   | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      | Parallel compare     | Good for parallel implementation                |

### Distribution Based

| Name                 | Best       | Average        | Worst      | Memory | Stable | In-place | Method         | Notes                             |
| -------------------- | ---------- | -------------- | ---------- | ------ | ------ | -------- | -------------- | --------------------------------- |
| **BlockSort**        | O(n log n) | O(n log n)     | O(n log n) | O(n)   | Yes    | No       | Block merging  | Cache-friendly merge sort variant |
| **BucketSort**       | O(n+k)     | O(n+k)         | O(n²)      | O(n+k) | Yes    | No       | Distribution   | Good for uniform distribution     |
| **CountingSort**     | O(n+k)     | O(n+k)         | O(n+k)     | O(k)   | Yes    | No       | Counting       | Excellent for small ranges        |
| **FlashSort**        | O(n)       | O(n)           | O(n²)      | O(n)   | No     | No       | Distribution   | Fast for uniform distribution     |
| **PigeonholeSort**   | O(n+r)     | O(n+r)         | O(n+r)     | O(r)   | Yes    | No       | Distribution   | r is the range of input           |
| **RadixSort**        | O(d(n+k))  | O(d(n+k))      | O(d(n+k))  | O(n+k) | Yes    | No       | Digit by digit | d is number of digits             |
| **SpreadSort**       | O(n)       | O(n log log n) | O(n log n) | O(n)   | No     | No       | Distribution   | Adaptive to input distribution    |
| **AmericanFlagSort** | O(n+k)     | O(n+k)         | O(n+k)     | O(1)   | Yes    | Yes      | MSD radix      | In-place radix sort variant       |
| **SampleSort**       | O(n log n) | O(n log n)     | O(n log n) | O(n)   | No     | No       | Distribution   | Super scalar sampling strategy    |

### Specialized

| Name               | Best       | Average    | Worst      | Memory | Stable | In-place | Method           | Notes                         |
| ------------------ | ---------- | ---------- | ---------- | ------ | ------ | -------- | ---------------- | ----------------------------- |
| **BurstSort**      | O(n log n) | O(n log n) | O(n log n) | O(n)   | Yes    | No       | Burst trie       | Originally for string sorting |
| **CycleSort**      | O(n²)      | O(n²)      | O(n²)      | O(1)   | No     | Yes      | Cycle detection  | Optimal memory writes         |
| **LibrarySort**    | O(n log n) | O(n log n) | O(n log n) | O(n)   | Yes    | No       | Gapped insertion | Also called Library sort      |
| **PatienceSort**   | O(n log n) | O(n log n) | O(n log n) | O(n)   | Yes    | No       | Card solitaire   | Based on patience card game   |
| **PostmanSort**    | O(n+r)     | O(n+r)     | O(n+r)     | O(n)   | Yes    | No       | Distribution     | r is the range of keys        |
| **SmoothSort**     | O(n)       | O(n log n) | O(n log n) | O(1)   | No     | Yes      | Leonardo heaps   | Variation of heapsort         |
| **StrandSort**     | O(n)       | O(n²)      | O(n²)      | O(n)   | Yes    | No       | Natural merging  | Good for nearly sorted data   |
| **TournamentSort** | O(n log n) | O(n log n) | O(n log n) | O(n)   | No     | No       | Tree-based       | Selection tree variant        |
| **TreeSort**       | O(n log n) | O(n log n) | O(n²)      | O(n)   | Yes    | No       | BST-based        | Binary search tree sort       |

### Modern Hybrids

| Name               | Best       | Average    | Worst      | Memory   | Stable | In-place | Method                        | Notes                                            |
| ------------------ | ---------- | ---------- | ---------- | -------- | ------ | -------- | ----------------------------- | ------------------------------------------------ |
| **GeneralSort**    | O(n)       | O(n log n) | O(n log n) | O(n)     | Yes    | No       | Adaptive hybrid               | Timsort-inspired adaptive sort                   |
| **HybridSort**     | O(n)       | O(n log n) | O(n log n) | O(n)     | Yes    | No       | Multi-strategy                | Combines counting/radix/insertion                |
| **UltimateSort**   | O(n log n) | O(n log n) | O(n log n) | O(log n) | No     | Yes      | Parallel hybrid               | Dual-pivot quick + parallel merge                |
| **CascadeSort**    | O(n)       | O(n log n) | O(n log n) | O(n)     | Yes    | No       | Cascading merge sort          | Efficient for partially sorted data              |
| **JupiterSort**    | O(n log n) | O(n log n) | O(n log n) | O(log n) | No     | Yes      | Parallel dual-pivot quicksort | Uses orbital partitioning and parallelism        |
| **WeaveMergeSort** | O(n log n) | O(n log n) | O(n log n) | O(n)     | Yes    | No       | Cache-optimized merge sort    | Improves cache locality with unrolled merge loop |
| **GallopingSort**  | O(n)       | O(n log n) | O(n log n) | O(1)     | Yes    | Yes      | Insertion sort with galloping | Excels on nearly sorted data                     |
| **QuantumSort**    | O(n log n) | O(n log n) | O(n log n) | O(n)     | Yes    | No       | Parallel merge sort           | Uses goroutines for concurrent sorting           |
| **GrailSort**      | O(n log n) | O(n log n) | O(n log n) | O(√n)    | Yes    | No       | Bottom-up merge sort          | Uses fixed buffer for efficient merging          |
| **CubeSort**       | O(n)       | O(n log n) | O(n log n) | O(n)     | Yes    | No       | Multi-way merge               | Uses cube root block size for merging            |
| **WikiSort**       | O(n log n) | O(n log n) | O(n log n) | O(1)     | Yes    | Yes      | Block merge sort              | In-place stable sorting algorithm                |

### Special Purpose

| Name            | Best | Average | Worst   | Memory  | Stable | In-place | Method          | Notes                                    |
| --------------- | ---- | ------- | ------- | ------- | ------ | -------- | --------------- | ---------------------------------------- |
| **BeadSort**    | O(n) | O(n\*m) | O(n\*m) | O(n\*m) | Yes    | No       | Gravity         | m is max value, visual/physical metaphor |
| **PancakeSort** | O(n) | O(n²)   | O(n²)   | O(1)    | No     | Yes      | Prefix reversal | Sorts by flipping prefixes of the array  |

### removed from benchmarking

- **BeadSort**

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

## Data Distributions

| Name                                                  | Description                               | Image                                                     |
| ----------------------------------------------------- | ----------------------------------------- | --------------------------------------------------------- |
| [AllZero](docs/distribution/AllZero.md)               | Array of all zeros                        | ![AllZero](images/distribution/AllZero.svg)               |
| [BackToFront](docs/distribution/BackToFront.md)       | Single element pushed from back to front  | ![BackToFront](images/distribution/BackToFront.svg)       |
| [FrontToBack](docs/distribution/FrontToBack.md)       | Single element pushed from front to back  | ![FrontToBack](images/distribution/FrontToBack.svg)       |
| [MiddleToBack](docs/distribution/MiddleToBack.md)     | Single element pushed from middle to back | ![MiddleToBack](images/distribution/MiddleToBack.svg)     |
| [Mountain](docs/distribution/Mountain.md)             | Elements form a mountain shape            | ![Mountain](images/distribution/Mountain.svg)             |
| [NearlyReversed](docs/distribution/NearlyReversed.md) | Almost completely reversed sequence       | ![NearlyReversed](images/distribution/NearlyReversed.svg) |
| [NearlySorted](docs/distribution/NearlySorted.md)     | Almost completely sorted sequence         | ![NearlySorted](images/distribution/NearlySorted.svg)     |
| [Plateau](docs/distribution/Plateau.md)               | Elements form a plateau shape             | ![Plateau](images/distribution/Plateau.svg)               |
| [PushMiddle](docs/distribution/PushMiddle.md)         | Elements pushed towards the middle        | ![PushMiddle](images/distribution/PushMiddle.svg)         |
| [Random](docs/distribution/Random.md)                 | Completely random distribution            | ![Random](images/distribution/Random.svg)                 |
| [RandomMaxN](docs/distribution/RandomMaxN.md)         | Random integers up to max N               | ![RandomMaxN](images/distribution/RandomMaxN.svg)         |
| [RandomMod8](docs/distribution/RandomMod8.md)         | Random values modulo 8                    | ![RandomMod8](images/distribution/RandomMod8.svg)         |
| [RandomMod16](docs/distribution/RandomMod16.md)       | Random values modulo 16                   | ![RandomMod16](images/distribution/RandomMod16.svg)       |
| [RepeatedMod8](docs/distribution/RepeatedMod8.md)     | Repeated sequence modulo 8                | ![RepeatedMod8](images/distribution/RepeatedMod8.svg)     |
| [RepeatedMod16](docs/distribution/RepeatedMod16.md)   | Repeated sequence modulo 16               | ![RepeatedMod16](images/distribution/RepeatedMod16.svg)   |
| [Reversed](docs/distribution/Reversed.md)             | Completely reversed sequence              | ![Reversed](images/distribution/Reversed.svg)             |
| [Rotated](docs/distribution/Rotated.md)               | Sorted sequence rotated by N positions    | ![Rotated](images/distribution/Rotated.svg)               |
| [SmallHills](docs/distribution/SmallHills.md)         | Elements form small hill patterns         | ![SmallHills](images/distribution/SmallHills.svg)         |
| [Sorted](docs/distribution/Sorted.md)                 | Completely sorted sequence                | ![Sorted](images/distribution/Sorted.svg)                 |
| [Valley](docs/distribution/Valley.md)                 | Elements form a valley shape              | ![Valley](images/distribution/Valley.svg)                 |

Usage License This project is licensed under the MIT License.
