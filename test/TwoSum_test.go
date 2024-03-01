package test

import (
	"leetcode/src"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	sum := src.TwoSum{}
	testCase := []struct {
		numbers  []int
		target   int
		expected []int
	}{{numbers: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}}}

	for _, tc := range testCase {
		result := sum.TwoSum(tc.numbers, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Actual : %v\nExpected: %v", result, tc.expected)
		}
	}
}
