package main

import "testing"

var arr, found, notFound = [6]int{4, 5, 6, 0, 11, 7}, "Found", "Not found"

func TestLinearSearchShouldFindTarget(t *testing.T) {
	targetInArr := 7
	got := linearSearch(arr, targetInArr)

	if got != found {
		t.Errorf("linearSearch(%v) == %v, expected %v", arr, got, found)
	}
}

func TestLinearSearchShouldNotFindTarget(t *testing.T) {
	targetInArr := 100
	got := linearSearch(arr, targetInArr)

	if got != notFound {
		t.Errorf("linearSearch(%v) == %v, expected %v", arr, got, notFound)
	}
}