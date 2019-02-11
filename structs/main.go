package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	rodolfo := person{"Rodolfo", "Leal"}
	alex := person{firstName: "Alex", lastName: "Anderson"}

	var joao person

	joao.firstName = "Joao"
	joao.lastName = "Silva"

	fmt.Println(rodolfo, alex, joao)
	fmt.Printf("%+v", alex)
}
