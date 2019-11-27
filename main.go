package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func checkAge(age int) bool {
	return age >= 18
}

func sellBeer(person Person) {
	if checkAge(person.Age) {
		fmt.Println(person.Name, "can buy beer")
	} else {
		fmt.Println(person.Name, "CAN'T buy beer")
	}
}

func main() {
	person := Person{
		Name: "Samuel",
		Age:  19,
	}
	sellBeer(person)
}
