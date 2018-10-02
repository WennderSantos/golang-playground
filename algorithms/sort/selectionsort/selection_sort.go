package main

import "fmt"

func main() {
	arr := []int {4, 10, 32, 1, 7, 8, 3, 6, 5}
	selectionSort(arr)

	fmt.Println(arr)
}

func getIndexWithMinimumValue(arr []int, fromIndex int) int {
	minimumValue := arr[fromIndex]
	indexWithMinimumValue := fromIndex

	for i := fromIndex + 1; i < len(arr); i++ {
		if arr[i] < minimumValue {
			indexWithMinimumValue = i
			minimumValue = arr[i]
		}
	}
	return indexWithMinimumValue
}

func swap(arr []int, firstIndex, secondIndex int) {
	aux := arr[secondIndex]
	arr[secondIndex] = arr[firstIndex]
	arr[firstIndex] = aux
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr) -1; i++ {
		indexToSwap := getIndexWithMinimumValue(arr, i)
		swap(arr, i, indexToSwap)
	}
}