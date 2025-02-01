package sortcomparison

import "math"

/*
CubeSort Implementation (3D Merge Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n)

Space Complexity:
  - O(n) - requires temporary storage for merging
  - Additional O(n^(1/3)) for group management

Implementation Notes:
  - Divides data into cube-like structure
  - Stable sort - maintains relative order of equal elements
  - Cache-efficient due to localized memory access
  - Parallelizable algorithm
  - Good for large datasets with sufficient memory
*/
func CubeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Calculate cube dimensions
	dim := int(math.Ceil(math.Pow(float64(n), 1.0/3.0)))
	if dim < 2 {
		dim = 2
	}

	// Create temporary storage
	temp := make([]int, n)
	copy(temp, arr)

	// Sort each dim x dim slice
	for i := 0; i < n; i += dim * dim {
		end := i + dim*dim
		if end > n {
			end = n
		}
		sortSlice(temp[i:end])
	}

	// Merge sorted slices
	mergeCube(arr, temp, dim, n)
}

func sortSlice(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func mergeCube(arr, temp []int, dim, n int) {
	// Merge in x dimension
	for z := 0; z < n; z += dim * dim {
		for y := 0; y < dim; y++ {
			mergeArrays(arr, temp, z+y*dim, z+(y+1)*dim-1, n)
		}
	}

	// Merge in y dimension
	for z := 0; z < n; z += dim * dim {
		end := z + dim*dim
		if end > n {
			end = n
		}
		mergeArrays(arr, temp, z, end-1, n)
	}

	// Final merge in z dimension
	mergeArrays(arr, temp, 0, n-1, n)
}

func mergeArrays(arr, temp []int, left, right, n int) {
	if right >= n {
		right = n - 1
	}
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	for i <= mid {
		arr[k] = temp[i]
		i++
		k++
	}

	for j <= right {
		arr[k] = temp[j]
		j++
		k++
	}

	// Update temp array for next merge
	copy(temp[left:right+1], arr[left:right+1])
}
