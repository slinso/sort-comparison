package sortcomparison

import (
	"math"
	"math/rand"
	"sort"
)

const (
	basecaseSize = 1024
	logBuckets   = 8
	numBuckets   = 1 << logBuckets
)

type bucket uint32

// Classifier handles the classification of elements into buckets
type Classifier struct {
	splitters    []int
	numSplitters int
	bktout       []bucket
	bktsize      []int
}

// SampleSort implements Super Scalar Sample Sort
func SampleSort(arr []int) {
	n := len(arr)
	if n < basecaseSize {
		sort.Ints(arr)
		return
	}

	// Draw and sort samples
	sampleSize := oversamplingFactor(n) * numBuckets
	samples := make([]int, sampleSize)
	drawSample(arr, samples)
	sort.Ints(samples)

	// Check if all samples are equal
	if samples[0] == samples[len(samples)-1] {
		sort.Ints(arr)
		return
	}

	// Classify elements
	bktout := make([]bucket, n)
	classifier := newClassifier(samples, bktout)
	classifier.classify(arr)

	// Distribute elements
	temp := make([]int, n)
	offset := 0
	prefixSum := make([]int, numBuckets)

	// Calculate prefix sums
	sum := 0
	for i := 0; i < numBuckets; i++ {
		prefixSum[i] = sum
		sum += classifier.bktsize[i]
	}

	// Distribute
	for i, v := range arr {
		pos := prefixSum[bktout[i]]
		temp[pos] = v
		prefixSum[bktout[i]]++
	}

	// Recursive sorting of buckets
	offset = 0
	for i := 0; i < numBuckets; i++ {
		size := classifier.bktsize[i]
		if size == 0 {
			continue
		}

		bucket := temp[offset : offset+size]
		if size <= basecaseSize || float64(n)/float64(size) < 2 {
			sort.Ints(bucket)
		} else {
			SampleSort(bucket)
		}
		copy(arr[offset:], bucket)
		offset += size
	}
}

func newClassifier(samples []int, bktout []bucket) *Classifier {
	c := &Classifier{
		splitters:    make([]int, 1<<logBuckets),
		numSplitters: (1 << logBuckets) - 1,
		bktout:       bktout,
		bktsize:      make([]int, 1<<logBuckets),
	}
	c.buildRecursive(samples, 0, len(samples)-1, 1)
	return c
}

func (c *Classifier) buildRecursive(samples []int, lo, hi, pos int) {
	mid := lo + (hi-lo)/2
	c.splitters[pos] = samples[mid]

	if 2*pos < c.numSplitters {
		c.buildRecursive(samples, lo, mid, 2*pos)
		c.buildRecursive(samples, mid+1, hi, 2*pos+1)
	}
}

func (c *Classifier) step(i bucket, key int) bucket {
	return 2*i + bucket(boolToInt(c.splitters[i] < key))
}

func (c *Classifier) findBucket(key int) bucket {
	i := bucket(1)
	for i <= bucket(c.numSplitters) {
		i = c.step(i, key)
	}
	return i - bucket(1<<logBuckets)
}

func (c *Classifier) classify(arr []int) {
	for i, v := range arr {
		bucket := c.findBucket(v)
		c.bktout[i] = bucket
		c.bktsize[bucket]++
	}
}

func oversamplingFactor(n int) int {
	r := math.Sqrt(float64(n) / float64(2*numBuckets*(logBuckets+4)))
	if r < 1 {
		return 1
	}
	return int(r)
}

func drawSample(arr []int, samples []int) {
	n := len(arr)
	for i := range samples {
		idx := rand.Intn(n)
		samples[i] = arr[idx]
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
