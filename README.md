h# Sort Comparison

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

| Name                 | Best       | Average     | Worst       | Memory   | Stable | In-place |
| -------------------- | ---------- | ----------- | ----------- | -------- | ------ | -------- |
| **FlashSort**        | O(n)       | O(n)        | O(n²)       | O(n)     | No     | No       |
| **AmericanFlagSort** | O(n+k)     | O(n+k)      | O(n+k)      | O(1)     | Yes    | Yes      |
| **CountingSort**     | O(n+k)     | O(n+k)      | O(n+k)      | O(k)     | Yes    | No       |
| **BucketSort**       | O(n+k)     | O(n+k)      | O(n²)       | O(n+k)   | Yes    | No       |
| **PigeonholeSort**   | O(n+r)     | O(n+r)      | O(n+r)      | O(r)     | Yes    | No       |
| **PostmanSort**      | O(n+r)     | O(n+r)      | O(n+r)      | O(n)     | Yes    | No       |
| **BlockSort**        | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **BurstSort**        | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **GrailSort**        | O(n log n) | O(n log n)  | O(n log n)  | O(√n)    | Yes    | No       |
| **HeapSort**         | O(n log n) | O(n log n)  | O(n log n)  | O(1)     | No     | Yes      |
| **IntroSort**        | O(n log n) | O(n log n)  | O(n log n)  | O(log n) | No     | Yes      |
| **JupiterSort**      | O(n log n) | O(n log n)  | O(n log n)  | O(log n) | No     | Yes      |
| **LibrarySort**      | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **MergeSort**        | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **PatienceSort**     | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **QuantumSort**      | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **SampleSort**       | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | No     | No       |
| **TournamentSort**   | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | No     | No       |
| **TreeSortAVL**      | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **UltimateSort**     | O(n log n) | O(n log n)  | O(n log n)  | O(log n) | No     | Yes      |
| **WeaveMergeSort**   | O(n log n) | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **WikiSort**         | O(n log n) | O(n log n)  | O(n log n)  | O(1)     | Yes    | Yes      |
| **CascadeSort**      | O(n)       | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **CubeSort**         | O(n)       | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **GallopingSort**    | O(n)       | O(n log n)  | O(n log n)  | O(1)     | Yes    | Yes      |
| **GeneralSort**      | O(n)       | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **HybridSort**       | O(n)       | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **SmoothSort**       | O(n)       | O(n log n)  | O(n log n)  | O(1)     | No     | Yes      |
| **TimSort**          | O(n)       | O(n log n)  | O(n log n)  | O(n)     | Yes    | No       |
| **CombSort**         | O(n log n) | O(n log n)  | O(n²)       | O(1)     | No     | Yes      |
| **QuickSort**        | O(n log n) | O(n log n)  | O(n²)       | O(log n) | No     | Yes      |
| **SpreadSort**       | O(n)       | O(n log² n) | O(n log² n) | O(n)     | No     | No       |
| **ShellSort**        | O(n log n) | O(n log² n) | O(n log² n) | O(1)     | No     | Yes      |
| **BubbleSort**       | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      |
| **CocktailSort**     | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      |
| **CycleSort**        | O(n²)      | O(n²)       | O(n²)       | O(1)     | No     | Yes      |
| **ExchangeSort**     | O(n²)      | O(n²)       | O(n²)       | O(1)     | No     | Yes      |
| **GnomeSort**        | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      |
| **InsertionSort**    | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      |
| **OddEvenSort**      | O(n)       | O(n²)       | O(n²)       | O(1)     | Yes    | Yes      |
| **PancakeSort**      | O(n)       | O(n²)       | O(n²)       | O(1)     | No     | Yes      |
| **SelectionSort**    | O(n²)      | O(n²)       | O(n²)       | O(1)     | No     | Yes      |
| **StrandSort**       | O(n)       | O(n²)       | O(n²)       | O(n)     | Yes    | No       |

### removed from benchmarking

| Name         | Best       | Average    | Worst   | Memory  | Stable | In-place | Notes                                                                         |
| ------------ | ---------- | ---------- | ------- | ------- | ------ | -------- | ----------------------------------------------------------------------------- |
| **BeadSort** | O(n)       | O(n\*m)    | O(n\*m) | O(n\*m) | Yes    | No       | visual/physical metaphor, BeadSortInspired with Bitset optimzations was added |
| **TreeSort** | O(n log n) | O(n log n) | O(n²)   | O(n)    | Yes    | No       | classic BST gets too slow for 10k elements                                    |

## Benchmarking

Attention: running all benchmarks really takes a long time. It is recommended to run benchmarks for specific algorithms or distributions.

```go
go test -bench=. -timeout 6h
```

Compare specific algorithms:

```go
go test -bench "//Quick/10"
```

Compare specific distributions:

```go
go test -bench "/RandomMax//10"
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
