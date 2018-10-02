package main

import ("testing"
				sliceUtil "github.com/wenndersantos/algorithms/sort/sliceutil")

func TestMergeSort(t *testing.T) {
	unsortedArr := []int{5, 11, 10, 3, 9, 10, 47, 0, 2, 1, 25}
	sortedArr := []int{0, 1, 2, 3, 5, 9, 10, 10, 11, 25, 47}

	got := MergeSort(unsortedArr)

	if !sliceUtil.Equal(got, sortedArr) {
		t.Errorf("MergeSort(%v) == %v, but expected %v", unsortedArr, got, sortedArr)
	}
}