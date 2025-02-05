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
	testSizes = []int{
		10,    // 10
		100,   // 100
		1000,  // 1K
		10000, // 10K
		// 100000,     // 100K
		// 1000000,    // 1M
		// 10000000,   // 10M
		// 100000000,  // 100M
		// 1000000000, // 1B
	}

	sortImplementations = []Sorter{
		{"StdSort", slices.Sort[[]int, int], nil},
		{"AdaptiveSort", nil, AdaptiveSort},
		{"AmericanFlagSort", AmericanFlagSort, nil},
		{"BeadSortInspired", nil, BeadSortInspired},
		{"BitonicSort", BitonicSort, nil},
		{"BitonicSortAny", BitonicSortAny, nil},
		{"BlockSort", BlockSort, nil},
		{"BubbleSortEarly", BubbleSortEarly[int], nil},
		{"BucketSort", BucketSort[int], nil},
		{"BurstSort", BurstSort, nil},
		{"CascadeSort", nil, CascadeSort},
		{"CocktailShakerSort", CocktailShakerSort, nil},
		{"CombSort", CombSort, nil},
		{"CountingSort", nil, CountingSort},
		{"CubeSort", CubeSort, nil},
		{"CycleSort", CycleSort, nil},
		{"ExchangeSort", ExchangeSort, nil},
		{"FlashSort", FlashSort, nil},
		{"GallopingSort", GallopingSort, nil},
		{"GnomeSort", GnomeSort, nil},
		{"GrailSort", GrailSort, nil},
		{"HeapSort", HeapSort, nil},
		{"HybridSortSonnet", nil, HybridSort},
		{"InsertionSort", InsertionSort, nil},
		{"IntroSort", IntroSort, nil},
		{"JupiterSort", JupiterSort, nil},
		{"LibrarySort", LibrarySort, nil},
		{"MergeSort", MergeSort, nil},
		{"OddEvenSort", nil, OddEvenSort},
		{"PancakeSort", PancakeSort, nil},
		{"PatienceSort", PatienceSort, nil},
		{"PigeonholeSort", PigeonholeSort, nil},
		{"PostmanSort", nil, PostmanSort},
		{"QuantumSortO3", nil, QuantumSort},
		{"QuickSort", QuickSort, nil},
		{"RadixSortLSD", RadixSortLSD, nil},
		{"RadixSortMSD", RadixSortMSD, nil},
		{"SampleSort", SampleSort, nil},
		{"SelectionSort", SelectionSort, nil},
		{"ShellSort", ShellSort, nil},
		{"SimplePancakeSort", SimplePancakeSort, nil},
		{"SmoothSort", SmoothSort, nil},
		{"SpreadSort", nil, SpreadSort},
		{"StrandSort", StrandSort, nil},
		{"TimSort", TimSort, nil},
		{"TournamentSort", TournamentSort, nil},
		{"TreeSort", TreeSort, nil},
		{"WeaveMergeSort", nil, WeaveMergeSort},
		{"WikiSort", nil, WikiSort},
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

// General benchmark function for sorting algorithms
func BenchmarkSort(b *testing.B) {
	for _, size := range testSizes {
		for _, gen := range dataGenerators {
			for _, s := range sortImplementations {
				if s.fn != nil && s.fnNotInPlace != nil {
					b.Fail()
				}

				if s.fn != nil {
					b.Run(fmt.Sprintf("dist=%s/algo=%s/size=%s", gen.name, s.name, formatSize(size)), func(b *testing.B) {
						data := gen.Data(size)
						testData := make([]int, len(data))

						for i := 0; i < b.N; i++ {
							copy(testData, data)

							s.fn(testData)
						}
					})
				}

				if s.fnNotInPlace != nil {
					b.Run(fmt.Sprintf("dist=%s/algo=%s/size=%s", gen.name, s.name, formatSize(size)), func(b *testing.B) {
						data := gen.Data(size)
						testData := make([]int, len(data))

						for i := 0; i < b.N; i++ {
							copy(testData, data)

							s.fnNotInPlace(testData)
						}
					})
				}
			}
		}
	}
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

	for _, s := range sortImplementations {
		for _, tc := range testCases {
			// special case for BitonicSort
			if s.name == "BitonicSort" {
				tc.size = 64
			}

			if s.fn != nil {
				t.Run(fmt.Sprintf("/Ret/%s/%s", s.name, tc.name), func(t *testing.T) {
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
