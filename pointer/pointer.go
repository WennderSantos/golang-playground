package main

import "fmt"

func main() {
	foo := "foo"
	fmt.Println(foo)
	aux(&foo)
        fmt.Println("Hello", foo)
}

func aux(foo *string) {
	*foo = "world"
}
