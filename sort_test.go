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
		{"BeadSort", nil, BeadSort},
		{"BitonicSort", nil, BitonicSort},
		{"BlockSort", BlockSort, nil},
		{"BubbleSort", BubbleSort, nil},
		{"BucketSort", BucketSort[int], nil},
		{"BurstSort", BurstSort, nil},
		{"CocktailShakerSort", CocktailShakerSort, nil},
		{"CombSort", CombSort, nil},
		{"CountingSort", CountingSort, CountingSortRet},
		{"CubeSort", CubeSort, nil},
		{"CycleSort", CycleSort, nil},
		{"ExchangeSort", ExchangeSort, nil},
		{"FlashSort", FlashSort, nil},
		{"GnomeSort", GnomeSort, nil},
		{"HeapSort", HeapSort, nil},
		{"InsertionSort", InsertionSort, nil},
		{"IntroSort", IntroSort, nil},
		{"LibrarySort", LibrarySort, nil},
		{"MergeSort", MergeSort, nil},
		{"OddEvenSort", OddEvenSort, nil},
		{"PatienceSort", PatienceSort, nil},
		{"PigeonholeSort", PigeonholeSort, nil},
		{"PostmanSort", nil, PostmanSort},
		{"QuickSort", QuickSort, nil},
		{"RadixSortLSD", RadixSort, nil},
		{"RadixSortMSD", RadixSortMSD, nil},
		// TODO: {"SampleSort", nil, SampleSort},
		{"SelectionSort", SelectionSort, nil},
		{"ShellSort", ShellSort, nil},
		{"SimplePancakeSort", SimplePancakeSort, nil},
		{"SmoothSort", SmoothSort, nil},
		{"SpreadSort", nil, SpreadSort},
		{"StrandSort", StrandSort, nil},
		{"TimSort", TimSort, nil},
		{"TournamentSort", TournamentSort, nil},
		{"TreeSort", TreeSort, nil},
	}

	dataGenerators = []DataGenerator{
		{"Random", generateRandomInts},
		{"RandomMaxN", generateRandomIntsMaxN},
		{"AllZero", generateAllZero},
		{"Sorted", generateSortedInts},
		{"Reversed", generateReversedInts},
		{"Mountain", generateMountain},
		{"Plateau", generatePlateau},
		{"Sawtooth", generateSawtooth},
		{"RepeatedMod8", generateRepeatedMod8},
		{"RepeatedMod16", generateRepeatedMod16},
		{"PushFront", generatePushFront},
		{"PushBack", generatePushBack},
		{"MiddleToBack", generateMiddleToBack},
		{"PushMiddle", generatePushMiddle},
		{"NearlySorted", generateNearlySorted},
		{"NearlyReversed", generateNearlyReversed},
		{"RandomMod8", generateRandomMod8},
		{"RandomMod16", generateRandomMod16},
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
	for i := 0; i < n/2; i++ {
		data[i] = i
	}
	for i := n / 2; i < n; i++ {
		data[i] = n - i
	}
	return data
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

func generateSawtooth(n int) []int {
	data := make([]int, n)
	for i := range data {
		if i%2 == 0 {
			data[i] = i / 2
		} else {
			data[i] = i
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

// General benchmark function for sorting algorithms
func BenchmarkSort(b *testing.B) {
	for _, size := range testSizes {
		for _, gen := range dataGenerators {
			for _, s := range sortImplementations {
				if s.fn != nil {
					b.Run(fmt.Sprintf("dist=%s/algo=%s/size=%s/typ=ip", gen.name, s.name, formatSize(size)), func(b *testing.B) {
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
					b.Run(fmt.Sprintf("dist=%s/algo=%s/size=%s/typ=ret", gen.name, s.name, formatSize(size)), func(b *testing.B) {
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
		{"Reversed_100", 100, generateReversedInts},
		{"Mountain_100", 100, generateMountain},
		{"Plateau_100", 100, generatePlateau},
		{"Sawtooth_100", 100, generateSawtooth},
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
