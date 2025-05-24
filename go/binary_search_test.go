package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},  // Target in the middle
		{[]int{1, 2, 3, 4, 5}, 1, 0},  // Target at the beginning
		{[]int{1, 2, 3, 4, 5}, 5, 4},  // Target at the end
		{[]int{1, 2, 3, 4, 5}, 6, -1}, // Target not found
		{[]int{}, 1, -1},              // Empty array
	}

	for _, test := range tests {
		result := BinarySearch(test.input, test.target)
		if result != test.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; expected %d", test.input, test.target, result, test.expected)
		}
	}
}

func TestBinarySearchLogN(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},  // Target in the middle
		{[]int{1, 2, 3, 4, 5}, 1, 0},  // Target at the beginning
		{[]int{1, 2, 3, 4, 5}, 5, 4},  // Target at the end
		{[]int{1, 2, 3, 4, 5}, 6, -1}, // Target not found
		{[]int{}, 1, -1},              // Empty array
	}

	for _, test := range tests {
		result := BinarySearchLogN(test.input, test.target)
		if result != test.expected {
			t.Errorf("BinarySearchLogN(%v, %d) = %d; expected %d", test.input, test.target, result, test.expected)
		}
	}
}
