package main

import (
	"fmt"
	"math/rand"
	"slices"
)

var sample = make([]int, 100)

// var sampleLit = []int{100, 200, 300, 1, 24, 5, 91, 2, 2, 2, 3}

var r = rand.New(rand.NewSource(100))

func main() {
	for i := range sample {
		sample[i] = r.Int()
	}
	SelectionSort(sample)
	fmt.Println(len(sample))
	fmt.Println(slices.IsSorted(sample))
}

func swap(in []int, firstIdx, secIdx int) {
	temp := in[firstIdx]
	in[firstIdx] = in[secIdx]
	in[secIdx] = temp
}

// Worst: O(n^2)
// Average: θ(n^2)
// Best: Ω(n^2)
// Space complexity: O(1)
// Sorted and unsorted parts of the array at the beginning the array
func SelectionSort(in []int) {
	// i is our sorted boundary
	for i := range in {
		lowestIdx := i
		// check in[i] against the rest of the slice
		for j := i + 1; j < len(in)-1; j++ {
			if in[j] < in[lowestIdx] {
				lowestIdx = j
			}
		}
		swap(in, lowestIdx, i)
	}
}

// Worst: O(n^2)
// Average: θ(n^2)
// Best: Ω(n)
// Space complexity: O(1)
// Biggest values bubble up to the end of the array. Can be optimized for sorting semi-sorted
// data by keeping track of how many non-swaps there were at the end of the array and retiring that many extra elements.
// Used in polyfill algorithms
func BubbleSort(in []int) {
	if slices.IsSorted(in) {
		return
	}
	for i := range in {
		swapped := false
		for j := 0; j < len(in)-i; j++ {
			if in[j] > in[j+1] {
				swapped = true
				swap(in, j, j+1)
			}
		}
		if !swapped {
			break
		}
	}
}
