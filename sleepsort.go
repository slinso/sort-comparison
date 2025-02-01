package sortcomparison

import (
	"time"
)

// SleepSort sorts a slice of non-negative integers using the Sleep Sort algorithm.
// SleepSort leverages goroutines and time.Sleep such that each number sleeps for a duration
// proportional to its value before being sent to a channel. Smaller numbers are sent earlier,
// resulting in a sorted order.
//
// For very large numbers, sleeping for a duration equal to the value can be inefficient.
// To mitigate this, a scaling factor is computed based on the maximum value so that the longest
// sleep is capped at a target duration (here 100ms). This keeps the algorithm GC friendly,
// efficient for demonstration, and idiomatic, although SleepSort remains a novelty rather than
// a practical sorting algorithm.
//
// Time Complexity:
//   - Average-case: O(n + (maxValue*multiplier))
//
// Space Complexity:
//   - O(n) for goroutines and channel buffering.
func SleepSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum value to determine scaling.
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	// Define a target maximum sleep duration.
	const targetMaxSleep = 100 * time.Millisecond
	// Compute multiplier so that the largest value sleeps for roughly targetMaxSleep.
	// If maxVal is 0, use a multiplier of 1.
	multiplier := time.Duration(1)
	if maxVal > 0 {
		multiplier = targetMaxSleep / time.Duration(maxVal)
		if multiplier < 1 {
			multiplier = 1
		}
	}

	resultChan := make(chan int, len(arr))

	// Launch a goroutine for each number.
	for _, v := range arr {
		go func(num int) {
			time.Sleep(time.Duration(num) * multiplier)
			resultChan <- num
		}(v)
	}

	// Collect the results as they arrive.
	result := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		result = append(result, <-resultChan)
	}
	close(resultChan)
	return result
}
