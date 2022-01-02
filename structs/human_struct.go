package structs

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//default constructor
func NewHuman() *Human {
	human := new(Human)
	return human
}

func NewHumanWithParameters(firstName, lastName string, age int) *Human {

	human := new(Human)
	human.FirstName = firstName
	human.LastName = lastName
	human.Age = age
	return human
}
