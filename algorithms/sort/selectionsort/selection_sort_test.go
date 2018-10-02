package main

import "testing"

func TestShouldBeAbleToGetTheIndexWithMinimumValueOfAnArray(t *testing.T) {
	minimumValue := 1
	startingOnIndex := 2
	indexWithMinimumValue := 6

	arr := []int {9, 0, 10, 15, 5, 6, minimumValue}
	got := getIndexWithMinimumValue(arr, startingOnIndex)

	if got != indexWithMinimumValue {
		t.Errorf("getMinimumValue(%v, %v) == %v, but expected %v", arr, startingOnIndex, got, indexWithMinimumValue)
	}
}

func TestShouldSwapTwoIndexes(t *testing.T) {
	arr := []int {4, 3, 10, 4}
	want := []int {3, 4, 10, 4}
	firstIndex := 0
	secondIndex := 1

	swap(arr, firstIndex, secondIndex)

	if !Equal (arr, want) {
		t.Errorf("swap(%v, %v, %v) == %v, but expected %v", arr, firstIndex, secondIndex, arr, want)
	}


}

func TestShouldSortAnSliceUsingSelectionSort(t *testing.T) {
	arr := []int {4, 3, 10, 5, 0}
	want := []int {0, 3, 4, 5, 10}

	selectionSort(arr)

	if !Equal (arr, want) {
		t.Errorf("selectionSort(%v) == %v, but expected %v", arr, arr, want)
	}


}