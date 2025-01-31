package sortcomparison

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

const seed = 42 // Fixed seed for reproducible tests

var (
	rnd *rand.Rand // Package-level random source
	// Test sizes from 10 to 1B, multiplying by 10 each time
	testSizes = []int{
		10, // 10
		// 100,  // 100
		// 1000, // 1K
		// 10000,      // 10K
		// 100000,     // 100K
		// 1000000,    // 1M
		// 10000000,   // 10M
		// 100000000,  // 100M
		// 1000000000, // 1B
	}
)

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

// Helper functions for data generation
func generateRandomInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = int(rnd.Int31())
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

// Common less function
func less(a, b int) bool {
	return a < b
}

// Benchmark functions for each sort
func BenchmarkQuickSort(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))
			for i := range data {
				log.Println(data[i])
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				QuickSort(testData, less)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				MergeSort(testData, less)
			}
		})
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				Heapsort(testData, less)
			}
		})
	}
}

func BenchmarkIntroSort(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				Introsort(testData, less)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				BubbleSort(testData, less)
			}
		})
	}
}

func BenchmarkTreeSort(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				TreeSort(testData, less)
			}
		})
	}
}

func BenchmarkBurstSort(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				BurstSort(testData, less, func(x int) int { return x })
			}
		})
	}
}

func BenchmarkRadixSortLSD(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				LSDRadixSort(testData)
			}
		})
	}
}

func BenchmarkRadixSortMSD(b *testing.B) {
	for _, size := range testSizes {
		b.Run(formatSize(size), func(b *testing.B) {
			data := generateRandomInts(size)
			testData := make([]int, len(data))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				copy(testData, data)
				MSDRadixSort(testData)
			}
		})
	}
}
