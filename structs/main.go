package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	rodolfo := person{"Rodolfo", "Leal"}
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(rodolfo, alex)
}
