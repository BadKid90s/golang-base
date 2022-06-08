package main

import (
	"fmt"
	"goworkspase/model"
)

func main() {
	person := model.NewPerson("zs")
	person.SetAge(18)
	person.SetSal(5000)

	println(person.GetAge())
	println(person.GetSal())
	fmt.Println(person.GetSal())
}
