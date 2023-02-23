package main_test

import (
	"testing"
)

func TestCompareArrays(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}

	if !CompareArrays(arr1, arr2) {
		t.Errorf("Test failed: die Arrays Sollten identisch sein")
	}
	arr2[2] = 10

	if CompareArrays(arr1, arr2) {
		t.Errorf("Test failed: Die beiden Arrays sollten unterschiedlich sein")
	}
}
func CompareArrays (a, b []int) bool{
	if len(a) != len (b) {
		return false
	}

	for i := 0; i < len(a); i++{
		if a[i] != b[i] {
			return false
		}
	}
	
return true
}

