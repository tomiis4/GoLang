package main

import (
	"testing"
	"reflect"
)

type TestCase struct {
	id       string
	input    []int
	expected []int
}

func TestSort(t *testing.T) {
	testCases := []TestCase {
		{
			id: "A",
			input: []int{1,2,3},
			expected: []int{1,2,3},
		},
		{
			id: "B",
			input: []int{4,-5,1,3},
			expected: []int{-5,1,3,4},
		},
		{
			id: "C",
			input: []int{5,5,6,2,3,1},
			expected: []int{1,2,3,5,5,6},
		},
		{
			id: "D",
			input: []int{-9,-4,-0,-2,8},
			expected: []int{-9,-4,-2,-0,8},
		},
	}

	for _, tCase := range testCases {
		id, input, expected := tCase.id, tCase.input, tCase.expected

		sort(input)

		if reflect.DeepEqual(input, expected) {
			t.Logf("Test %v has passed succesfully.", id)
		} else {
			t.Errorf("Failed to sort array with ID: %v. Expected output: %v, got %v", id, expected, input)
		}
	}
}
