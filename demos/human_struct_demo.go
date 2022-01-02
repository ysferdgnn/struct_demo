package demos

import (
	"fmt"
	"struct_demo/structs"
)

func InitHumanStructNamedParameters() {

	fmt.Println("\n Init human struct with named parameters..")
	human := structs.Human{FirstName: "Yusuf", LastName: "Erdoğan", Age: 27}

	fmt.Println(human)
}

func InitHumanStructNamedParametersOnlyName() {
	fmt.Println("\n Init human struct with only name parameter..")

	human := structs.Human{FirstName: "Yusuf"}

	fmt.Println(human)
}

func InitHumanStructWithNewFunction() {
	fmt.Println("\n Init human struct with new function..")

	human := new(structs.Human)
	human.FirstName = "Yusuf"
	human.LastName = "Erdoğan"
	human.Age = 27
	fmt.Println(human)
}

func InitHumanStructWithConstructorFunction() {
	fmt.Println("\n Init human struct with constructor function..")

	human := *structs.NewHuman()
	human.FirstName = "Yusuf"
	human.LastName = "Erdoğan"
	human.Age = 27

	fmt.Println(human)
}

func InitHumanStructWithParameterizedConstructor() {

	fmt.Println("\n Init human struct with parameterized constructor function")

	human := structs.NewHumanWithParameters("Yusuf", "Erdoğan", 27)

	fmt.Println(human)
}
