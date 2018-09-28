package binarysearch

import "fmt"

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10

	fmt.Println(BinarySearch(arr, target))
}

func BinarySearch(arr [10]int, target int) string {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return "Found"
		} else if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		}
	}

	return "Not found"
}
