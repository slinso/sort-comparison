package sortcomparison

import "math/rand"

// random initialization seed
const seed = 42 // Fixed seed for reproducible tests
var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(seed))
}

// Data generation helpers
func GenerateRandomInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = int(rnd.Int31())
	}
	return data
}

// Data generation helpers
func GenerateRandomIntsMaxN(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rnd.Intn(n)
	}
	return data
}

func GenerateAllZero(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = 0
	}
	return data
}

func GenerateSortedInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	return data
}

func GenerateReversedInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = n - i
	}
	return data
}

func GenerateMountain(n int) []int {
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

func GenerateRotated(n int) []int {
	data := GenerateSortedInts(n)
	pivot := n / 2
	return append(data[pivot:], data[:pivot]...)
}

func GeneratePlateau(n int) []int {
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
func GenerateMod8(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i % 8
	}
	return data
}

// add a sawtooth distribution implementation
func GenerateMod16(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i % 16
	}
	return data
}

// add a small hills distribution implementation
func GenerateSmallHills(n int) []int {
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

func GenerateRepeatedMod8(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i % 8
	}
	return data
}

func GenerateRepeatedMod16(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i % 16
	}
	return data
}

func GenerateBackToFront(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i + 1
	}
	data[n-1] = 0

	return data
}

func GenerateFrontToBack(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	data[0] = n

	return data
}

func GenerateMiddleToBack(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	data[n/2] = n

	return data
}

func GeneratePushMiddle(n int) []int {
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

func GenerateNearlySorted(n int) []int {
	data := GenerateSortedInts(n)
	shufflePercentage(data, 5)

	return data
}

func GenerateNearlyReversed(n int) []int {
	data := GenerateReversedInts(n)
	shufflePercentage(data, 5)

	return data
}

func GenerateRandomMod8(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rnd.Intn(8)
	}
	return data
}

func GenerateRandomMod16(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rnd.Intn(16)
	}
	return data
}

func GenerateValley(n int) []int {
	data := make([]int, n)
	for i := 0; i < n/2; i++ {
		data[i] = n/2 - i
	}
	for i := n / 2; i < n; i++ {
		data[i] = i - n/2
	}
	return data
}
