package person

type Person struct {
        FirstName string
        LastName string
}

func ChangeName(person *Person) {
        person.LastName = "dos Santos"
}
