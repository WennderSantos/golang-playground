package binarysearch

import "testing"

var arr, found, notFound = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, "Found", "Not found"

func TestBinarySearchShouldFindTarget(t *testing.T) {
	targetInArr := 8

	got := BinarySearch(arr, targetInArr)
	if got != found {
		t.Errorf("BinarySearch(%q, %q) == %q, but want %q", arr, targetInArr, got, found)
	}
}

func TestBinarySearchShouldNotFindTarget(t *testing.T) {
	targetNotInArr := 19

	got := BinarySearch(arr, targetNotInArr)
	if got != notFound {
		t.Errorf("BinarySearch(%q, %q) == %q, but want %q", arr, targetNotInArr, got, notFound)
	}
}

