package main

import "fmt"

func main() {
	arr := [6]int{3, 5, 2, 1, 0, 9}
	target := 4
	fmt.Println(linearSearch(arr, target))
}

func linearSearch(arr [6]int, target int) string {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return "Found"
		}
	}
	return "Not found"
}
