package main

import (
		"fmt"
		"github.com/wenndersantos/struct-pointer/person"
       )

func main() {
	wennder := person.Person {
		FirstName: "Wennder",
		LastName: "xpto",
	} 	

	fmt.Println(wennder)	
	person.ChangeName(&wennder)
	fmt.Println(wennder)
}
