package main

import (
	"testing"
      	"github.com/wenndersantos/struct-pointer/person"
       )

func TestChangeName(t *testing.T) {
	wennder := person.Person {
		FirstName: "Wennder",
		LastName: "xpto",
	}

	person.ChangeName(&wennder)
	lastNameWanted := "dos Santos"
	lastNameGot := wennder.LastName

	if lastNameWanted != lastNameGot {
		t.Errorf("lastNameWanted == %q, got %q", lastNameWanted, lastNameGot)
	}
}
