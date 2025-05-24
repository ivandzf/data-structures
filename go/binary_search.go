package main

// sorted array integers
// find target integers
// if found return it index, else -1
// should run log(n)
func BinarySearch(input []int, target int) int {
	idx := -1
	for i := range input {
		if input[i] == target {
			idx = i
		}
	}

	return idx
}

func BinarySearchLogN(input []int, target int) int {
	idx := -1
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := low + (high-low)/2
		midValue := input[mid]

		if midValue == target {
			return mid
		} else if midValue < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return idx
}
