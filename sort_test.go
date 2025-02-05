package sortcomparison

import (
	"fmt"
	"slices"
	"testing"
)

// SortFunc represents a sorting function implementation
type SortFunc func([]int)

type SortNotInPlaceFunc func([]int) []int

type DataGenerator struct {
	name string
	Data func(int) []int
}

type Sorter struct {
	name         string
	fn           SortFunc
	fnNotInPlace SortNotInPlaceFunc
}

var (
	smallTestSizes = []int{
		10,    // 10
		100,   // 100
		1000,  // 1K
		10000, // 10K
	}

	mediumTestSizes = []int{
		100000,  // 100K
		1000000, // 1M
	}

	bigTestSizes = []int{
		10000000,   // 10M
		100000000,  // 100M
		1000000000, // 1B
	}

	testOnlyImplementations = []Sorter{
		{"BeadSort", BeadSort, nil},
	}

	bigSortImplementations = []Sorter{
		{"StdSort", slices.Sort[[]int, int], nil},
		{"AdaptiveSort", nil, AdaptiveSort},
		{"AmericanFlagSort", AmericanFlagSort, nil},
		{"BurstSort", BurstSort, nil},
		{"CascadeSort", nil, CascadeSort},
		{"CombSort", CombSort, nil},
		{"CountingSort", CountingSort, nil},
		{"FlashSort", FlashSort, nil},
		{"GrailSort", GrailSort, nil},
		{"HeapSort", HeapSort, nil},
		{"HybridSortSonnet", HybridSort, nil},
		{"IntroSort", IntroSort, nil},
		{"JupiterSort", JupiterSort, nil},
		{"MergeSort", MergeSort, nil},
		{"PigeonholeSort", PigeonholeSort, nil},
		{"PostmanSort", nil, PostmanSort},
		{"QuickSort", QuickSort, nil},
		{"RadixSortLSD", RadixSortLSD, nil},
		{"RadixSortMSD", RadixSortMSD, nil},
		{"SampleSort", SampleSort, nil},
		{"SpreadSort", nil, SpreadSort},
		{"TimSort", TimSort, nil},
		{"TournamentSort", TournamentSort, nil},
		{"WeaveMergeSort", nil, WeaveMergeSort},
	}

	mediumSortImplementations = []Sorter{
		{"BitonicSort", BitonicSort, nil},
		{"BitonicSortAny", BitonicSortAny, nil},
		{"BlockSort", BlockSort, nil},
		{"CubeSort", CubeSort, nil},
		{"PatienceSort", PatienceSort, nil},
		{"ShellSort", ShellSort, nil},
		{"StrandSort", StrandSort, nil},
		{"TreeSortAVL", TreeSortAVL, nil},
		{"WikiSort", nil, WikiSort},
	}

	smallSortImplementations = []Sorter{
		{"BeadSortInspired", nil, BeadSortInspired},
		{"BubbleSortEarly", BubbleSortEarly[int], nil},
		{"BucketSort", BucketSort[int], nil},
		{"CocktailShakerSort", CocktailShakerSort, nil},
		{"CycleSortOpt", CycleSortOpt, nil},
		{"ExchangeSort", ExchangeSort, nil},
		{"GallopingSort", GallopingSort, nil},
		{"GnomeSort", GnomeSort, nil},
		{"InsertionSort", InsertionSort, nil},
		{"LibrarySort", LibrarySort, nil},
		{"OddEvenSort", nil, OddEvenSort},
		{"PancakeSort", PancakeSort, nil},
		{"SelectionSort", SelectionSort, nil},
		{"SmoothSort", SmoothSort, nil},
		{"TreeSort", TreeSort, nil},
	}

	dataGenerators = []DataGenerator{
		{"Random", GenerateRandomInts},
		{"RandomMaxN", GenerateRandomIntsMaxN},
		{"AllZero", GenerateAllZero},
		{"Sorted", GenerateSortedInts},
		{"Rotated", GenerateRotated},
		{"Reversed", GenerateReversedInts},
		{"Mountain", GenerateMountain},
		{"Valley", GenerateValley},
		{"Plateau", GeneratePlateau},
		{"SmallHills", GenerateSmallHills},
		{"RandomMod8", GenerateRandomMod8},
		{"RepeatedMod8", GenerateRepeatedMod8},
		{"RandomMod16", GenerateRandomMod16},
		{"RepeatedMod16", GenerateRepeatedMod16},
		{"BackToFront", GenerateBackToFront},
		{"FrontToBack", GenerateFrontToBack},
		{"MiddleToBack", GenerateMiddleToBack},
		{"PushMiddle", GeneratePushMiddle},
		{"NearlySorted", GenerateNearlySorted},
		{"NearlyReversed", GenerateNearlyReversed},
	}
)

// Helper function to format size for benchmark name
func formatSize(size int) string {
	switch {
	// case size >= 1000000000:
	// 	return fmt.Sprintf("%dB", size/1000000000)
	// case size >= 1000000:
	// 	return fmt.Sprintf("%dM", size/1000000)
	// case size >= 1000:
	// 	return fmt.Sprintf("%dK", size/1000)
	default:
		return fmt.Sprintf("%d", size)
	}
}

// Helper closure to eliminate duplicate benchmarking code.
func runBenchmark(b *testing.B, gen DataGenerator, sortName string, size int, sortFunc func([]int)) {
	benchmarkName := fmt.Sprintf("dist=%s/algo=%s/size=%d", gen.name, sortName, size)
	b.Run(benchmarkName, func(b *testing.B) {
		data := gen.Data(size)
		testData := make([]int, len(data))

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// copy data to start with the same initial state every time
			copy(testData, data)
			sortFunc(testData)
		}
	})
}

func runSorterSuite(b *testing.B, sizes []int, sorter []Sorter) {
	for _, size := range sizes {
		for _, gen := range dataGenerators {
			for _, s := range sorter {
				if s.fn != nil && s.fnNotInPlace != nil {
					b.Fatalf("Sorter %s: define either fn or fnNotInPlace, not both", s.name)
				}

				if s.fn != nil {
					runBenchmark(b, gen, s.name, size, s.fn)
				}
				if s.fnNotInPlace != nil {
					runBenchmark(b, gen, s.name, size, func(data []int) {
						_ = s.fnNotInPlace(data)
					})
				}
			}
		}
	}
}

// General benchmark function for sorting algorithms
func BenchmarkSort(b *testing.B) {
	runSorterSuite(b, smallTestSizes, slices.Concat(smallSortImplementations, mediumSortImplementations, bigSortImplementations))
	runSorterSuite(b, mediumTestSizes, slices.Concat(mediumSortImplementations, bigSortImplementations))
	runSorterSuite(b, bigTestSizes, bigSortImplementations)
}

// testSortImplementation verifies a sorting implementation
func TestSort(t *testing.T) {
	testCases := []struct {
		name string
		size int
		gen  func(int) []int
	}{
		{"Random_100", 100, GenerateRandomIntsMaxN},
		{"AllZero_100", 100, GenerateAllZero},
		{"Sorted_100", 100, GenerateSortedInts},
		{"Rotated_100", 100, GenerateRotated},
		{"Reversed_100", 100, GenerateReversedInts},
		{"Mountain_100", 100, GenerateMountain},
		{"Valley_100", 100, GenerateValley},
		{"Plateau_100", 100, GeneratePlateau},
		{"SmallHills_100", 100, GenerateSmallHills},
		{"RepeatedMod8_100", 100, GenerateRepeatedMod8},
		{"RepeatedMod16_100", 100, GenerateRepeatedMod16},
		{"PushFront_100", 100, GenerateBackToFront},
		{"PushBack_100", 100, GenerateFrontToBack},
		{"MiddleToBack_100", 100, GenerateMiddleToBack},
		{"PushMiddle_100", 100, GeneratePushMiddle},
		{"NearlySorted_100", 100, GenerateNearlySorted},
		{"NearlyReversed_100", 100, GenerateNearlyReversed},
		{"RandomMod8_100", 100, GenerateRandomMod8},
		{"RandomMod16_100", 100, GenerateRandomMod16},
	}

	sorter := slices.Concat(testOnlyImplementations, smallSortImplementations, mediumSortImplementations, bigSortImplementations)

	for _, s := range sorter {
		for _, tc := range testCases {
			// special case for BitonicSort
			if s.name == "BitonicSort" {
				tc.size = 64
			}

			if s.fn != nil {
				t.Run(fmt.Sprintf("%s/%s", s.name, tc.name), func(t *testing.T) {
					data := tc.gen(tc.size)
					s.fn(data)

					// assert size
					if len(data) != tc.size {
						t.Errorf("Array size changed")
					}

					// assert that the array is sorted
					if !slices.IsSorted(data) {
						t.Errorf("Array not sorted: %v", data)
					}
				})
			}

			if s.fnNotInPlace != nil {
				t.Run(fmt.Sprintf("%s/%s", s.name, tc.name), func(t *testing.T) {
					data := tc.gen(tc.size)
					data = s.fnNotInPlace(data)

					// assert size
					if len(data) != tc.size {
						t.Errorf("Array size changed")
					}

					// assert that the array is sorted
					if !slices.IsSorted(data) {
						t.Errorf("Array not sorted: %v", data)
					}
				})
			}
		}
	}
}
