package sortcomparison

/*
PostmanSort Implementation (Route-Based Distribution Sort)

Time Complexity:
  - Average: O(n + r) where r is number of routes
  - Worst:   O(n²) when many elements in same route
  - Best:    O(n) when elements evenly distributed

Space Complexity:
  - O(n) for route storage
  - O(r) for route management where r is route count

Implementation Notes:
  - Similar to bucket sort but with dynamic route sizing
  - Stable sort - maintains relative order of equal elements
  - Efficient for data with natural clustering
  - Adaptive to data distribution patterns
*/
func PostmanSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find range for route distribution
	min, max := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	// Calculate optimal number of routes
	routeSize := (max-min)/n + 1
	if routeSize < 1 {
		routeSize = 1
	}
	numRoutes := (max-min)/routeSize + 1

	// Create routes
	routes := make([][]int, numRoutes)
	for i := range routes {
		routes[i] = make([]int, 0)
	}

	// Distribute elements to routes
	for _, val := range arr {
		routeIndex := (val - min) / routeSize
		routes[routeIndex] = append(routes[routeIndex], val)
	}

	// Sort each route using insertion sort
	for i := range routes {
		insertionSortRoute(routes[i])
	}

	// Merge routes back to original array
	index := 0
	for _, route := range routes {
		for _, val := range route {
			arr[index] = val
			index++
		}
	}
}

func insertionSortRoute(route []int) {
	for i := 1; i < len(route); i++ {
		key := route[i]
		j := i - 1
		for j >= 0 && route[j] > key {
			route[j+1] = route[j]
			j--
		}
		route[j+1] = key
	}
}
