package sortcomparison

import (
	"fmt"
	"math/rand"
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
	rnd       *rand.Rand
	testSizes = []int{
		10,         // 10
		100,        // 100
		1000,       // 1K
		10000,      // 10K
		100000,     // 100K
		1000000,    // 1M
		10000000,   // 10M
		100000000,  // 100M
		1000000000, // 1B
	}

	sortImplementations = []Sorter{
		{"StdSort", slices.Sort[[]int, int], nil},
		{"AdaptiveSort", nil, AdaptiveSort},
		{"AmericanFlagSort", AmericanFlagSort, nil},
		{"BeadSort", BeadSort, nil},
		{"BeadSortInspired", nil, BeadSortInspired},
		{"BitonicSort", BitonicSort, nil},
		{"BitonicSortAny", BitonicSortAny, nil},
		{"BlockSort", BlockSort, nil},
		{"BubbleSort", BubbleSort, nil},
		{"BucketSort", BucketSort[int], nil},
		{"BurstSort", BurstSort, nil},
		{"CascadeSort", nil, CascadeSort},
		{"CocktailShakerSort", CocktailShakerSort, nil},
		{"CombSort", CombSort, nil},
		{"CountingSort", CountingSort, CountingSortRet},
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
		{"QuantumSort", nil, QuantumSort},
		{"PostmanSort", nil, PostmanSort},
		{"QuickSort", QuickSort, nil},
		{"RadixSortLSD", RadixSort, nil},
		{"RadixSortMSD", RadixSortMSD, nil},
		// {"SampleSort", nil, SampleSort},
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
		{"Random", generateRandomInts},
		{"RandomMaxN", generateRandomIntsMaxN},
		{"AllZero", generateAllZero},
		{"Sorted", generateSortedInts},
		{"Rotated", generateRotated},
		{"Reversed", generateReversedInts},
		{"Mountain", generateMountain},
		{"Valley", generateValley},
		{"Plateau", generatePlateau},
		{"SmallHills", generateSmallHills},
		{"RandomMod8", generateRandomMod8},
		{"RepeatedMod8", generateRepeatedMod8},
		{"RandomMod16", generateRandomMod16},
		{"RepeatedMod16", generateRepeatedMod16},
		{"PushFront", generatePushFront},
		{"PushBack", generatePushBack},
		{"MiddleToBack", generateMiddleToBack},
		{"PushMiddle", generatePushMiddle},
		{"NearlySorted", generateNearlySorted},
		{"NearlyReversed", generateNearlyReversed},
	}
)

// random initialization seed
const seed = 42 // Fixed seed for reproducible tests

func init() {
	rnd = rand.New(rand.NewSource(seed))
}

// Helper function to format size for benchmark name
func formatSize(size int) string {
	switch {
	case size >= 1000000000:
		return fmt.Sprintf("%dB", size/1000000000)
	case size >= 1000000:
		return fmt.Sprintf("%dM", size/1000000)
	case size >= 1000:
		return fmt.Sprintf("%dK", size/1000)
	default:
		return fmt.Sprintf("%d", size)
	}
}

// Data generation helpers
func generateRandomInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = int(rnd.Int31())
	}
	return data
}

// Data generation helpers
func generateRandomIntsMaxN(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rnd.Intn(n)
	}
	return data
}

func generateAllZero(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = 0
	}
	return data
}

func generateSortedInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	return data
}

func generateReversedInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = n - i
	}
	return data
}

func generateMountain(n int) []int {
	data := make([]int, n)
	mid := n / 2
	for i := 0; i < n; i++ {
		if i <= mid {
			data[i] = i
		} else {
			data[i] = n - i
		}
	}
	return data
}

func generateRotated(n int) []int {
	data := generateSortedInts(n)
	pivot := n / 2
	return append(data[pivot:], data[:pivot]...)
}

func generatePlateau(n int) []int {
	data := make([]int, n)
	for i := 0; i < n/2; i++ {
		data[i] = i
	}
	for i := n / 2; i < n; i++ {
		data[i] = n / 2
	}
	return data
}

// add a sawtooth distribution implementation
func generateMod8(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i % 8
	}
	return data
}

// add a sawtooth distribution implementation
func generateMod16(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i % 16
	}
	return data
}

// add a small hills distribution implementation
func generateSmallHills(n int) []int {
	data := make([]int, n)
	segment := 20
	if n < segment {
		segment = n
	}
	for i := 0; i < n; i += segment {
		end := i + segment
		if end > n {
			end = n
		}
		mid := (i + end) / 2
		for j := i; j < end; j++ {
			if j <= mid {
				data[j] = j
			} else {
				data[j] = end - j + i
			}
		}
	}
	return data
}

func generateRepeatedMod8(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i % 8
	}
	return data
}

func generateRepeatedMod16(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i % 16
	}
	return data
}

func generatePushFront(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i + 1
	}
	data[n-1] = 0

	return data
}

func generatePushBack(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	data[0] = n

	return data
}

func generateMiddleToBack(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	data[n/2] = n

	return data
}

func generatePushMiddle(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	data[n-1] = n / 2

	return data
}

// shufflePercentage randomly swaps elements in the input slice based on the given percentage.
// Each swap exchanges two randomly selected elements in the slice.
//
// Parameters:
//   - data: slice of integers to be shuffled
//   - percentage: percentage of elements to be shuffled (0-100)
//
// The function modifies the input slice in place.
func shufflePercentage(data []int, percentage int) {
	n := len(data)

	if percentage < 0 || percentage > 100 {
		panic("invalid percentage")
	}

	for i := 0; i < n*percentage/100; i++ {
		a, b := rand.Intn(n), rand.Intn(n)
		data[a], data[b] = data[b], data[a]
	}
}

func generateNearlySorted(n int) []int {
	data := generateSortedInts(n)
	shufflePercentage(data, 5)

	return data
}

func generateNearlyReversed(n int) []int {
	data := generateReversedInts(n)
	shufflePercentage(data, 5)

	return data
}

func generateRandomMod8(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rnd.Intn(8)
	}
	return data
}

func generateRandomMod16(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rnd.Intn(16)
	}
	return data
}

func generateValley(n int) []int {
	data := make([]int, n)
	for i := 0; i < n/2; i++ {
		data[i] = n/2 - i
	}
	for i := n / 2; i < n; i++ {
		data[i] = i - n/2
	}
	return data
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
						b.ResetTimer()
						b.StopTimer()
						for i := 0; i < b.N; i++ {
							testData := make([]int, len(data))
							copy(testData, data)

							b.StartTimer()
							s.fn(testData)
							b.StopTimer()
						}
					})
				}

				if s.fnNotInPlace != nil {
					b.Run(fmt.Sprintf("dist=%s/algo=%s/size=%s", gen.name, s.name, formatSize(size)), func(b *testing.B) {
						data := gen.Data(size)
						b.ResetTimer()
						b.StopTimer()
						for i := 0; i < b.N; i++ {
							testData := make([]int, len(data))
							copy(testData, data)

							b.StartTimer()
							s.fnNotInPlace(testData)
							b.StopTimer()
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
		{"Random_100", 100, generateRandomIntsMaxN},
		{"AllZero_100", 100, generateAllZero},
		{"Sorted_100", 100, generateSortedInts},
		{"Rotated_100", 100, generateRotated},
		{"Reversed_100", 100, generateReversedInts},
		{"Mountain_100", 100, generateMountain},
		{"Valley_100", 100, generateValley},
		{"Plateau_100", 100, generatePlateau},
		{"SmallHills_100", 100, generateSmallHills},
		{"RepeatedMod8_100", 100, generateRepeatedMod8},
		{"RepeatedMod16_100", 100, generateRepeatedMod16},
		{"PushFront_100", 100, generatePushFront},
		{"PushBack_100", 100, generatePushBack},
		{"MiddleToBack_100", 100, generateMiddleToBack},
		{"PushMiddle_100", 100, generatePushMiddle},
		{"NearlySorted_100", 100, generateNearlySorted},
		{"NearlyReversed_100", 100, generateNearlyReversed},
		{"RandomMod8_100", 100, generateRandomMod8},
		{"RandomMod16_100", 100, generateRandomMod16},
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
