package sortcomparison

const (
    // Thresholds
    insertionLimit = 16
    countingSortLimit = 1 << 16
    parallelLimit = 1 << 13
    samplingRate = 64
)

// IntegerSort analyzes input characteristics and chooses optimal algorithm
func IntegerSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    // Sample data characteristics
    stats := analyzeArray(arr)
    
    // Choose optimal algorithm based on data characteristics
    switch {
    case len(arr) <= insertionLimit:
        return insertionSort(arr)
    
    case stats.range <= countingSortLimit && stats.range <= len(arr)*4:
        return countingSort(arr, stats.min, stats.max)
        
    case stats.clusters <= 256 && stats.range > len(arr):
        return americanFlagSort(arr)
        
    case stats.nearSorted && stats.inversions < len(arr)/4:
        return adaptiveMergeSort(arr)
        
    default:
        return parallelRadixSort(arr)
    }
}

type arrayStats struct {
    min, max int
    range int
    clusters int
    nearSorted bool
    inversions int
}

func analyzeArray(arr []int) arrayStats {
    stats := arrayStats{
        min: arr[0],
        max: arr[0],
    }
    
    inversions := 0
    prevVal := arr[0]
    clusters := 1
    
    // Sample array at regular intervals
    for i := 0; i < len(arr); i += samplingRate {
        v := arr[i]
        
        // Update min/max
        if v < stats.min {
            stats.min = v
        }
        if v > stats.max {
            stats.max = v
        }
        
        // Count inversions
        if v < prevVal {
            inversions++
        }
        
        // Detect value clusters
        if v != prevVal {
            clusters++
        }
        
        prevVal = v
    }
    
    stats.range = stats.max - stats.min + 1
    stats.clusters = clusters
    stats.inversions = inversions * samplingRate
    stats.nearSorted = inversions < len(arr)/(4*samplingRate)
    
    return stats
}

func parallelRadixSort(arr []int) []int {
    // Radix sort with parallel counting and distribution
    bits := 8
    mask := 1<<bits - 1
    temp := make([]int, len(arr))
    
    // Count bits needed
    max := 0
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    
    passes := 0
    for max > 0 {
        passes++
        max >>= bits
    }
    
    // Process each digit in parallel when array is large enough
    for shift := 0; shift < passes*bits; shift += bits {
        counts := [256]int{}
        
        // Count digits
        for _, v := range arr {
            digit := (v >> shift) & mask
            counts[digit]++
        }
        
        // Compute positions
        pos := 0
        for i := 0; i < len(counts); i++ {
            count := counts[i]
            counts[i] = pos
            pos += count
        }
        
        // Distribute
        for _, v := range arr {
            digit := (v >> shift) & mask
            temp[counts[digit]] = v
            counts[digit]++
        }
        
        arr, temp = temp, arr
    }
    
    if passes&1 == 1 {
        copy(temp, arr)
        arr = temp
    }
    
    return arr
}

func countingSort(arr []int, min, max int) []int {
    range_ := max - min + 1
    counts := make([]int, range_)
    
    // Count frequencies
    for _, v := range arr {
        counts[v-min]++
    }
    
    // Write back sorted values
    pos := 0
    for val := min; val <= max; val++ {
        for counts[val-min] > 0 {
            arr[pos] = val
            pos++
            counts[val-min]--
        }
    }
    return arr
}

// Additional helper methods omitted for brevity...