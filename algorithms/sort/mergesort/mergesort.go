package main

import (
	"fmt"
)

func main() {
	unsortedArr := []int{5, 11, 10, 3, 9, 47, 0, 2, 1, 25}

	sortedArr := MergeSort(unsortedArr)
	fmt.Println(sortedArr)
}

func MergeSort(arr []int) []int {
	arrLenght := len(arr)

	if arrLenght == 1 {
		return arr
	}

	middle := int64(arrLenght / 2)
	leftSide := arr[:middle]
	rightSide := arr[middle:]

	return Merge(MergeSort(leftSide), MergeSort(rightSide))
}

func Merge(leftSide []int, rightSide []int) []int {
	leftIndex := 0
	rightIndex := 0
	orderedArr := make([]int, 0)

	for (leftIndex < len(leftSide) && rightIndex < len(rightSide)) {
		if leftSide[leftIndex] < rightSide[rightIndex] {
			orderedArr = append(orderedArr, leftSide[leftIndex])
			leftIndex++
		} else {
			orderedArr = append(orderedArr, rightSide[rightIndex])
			rightIndex++
		}
	}

	return append(orderedArr, append(leftSide[leftIndex:], rightSide[rightIndex:]...)...)
}