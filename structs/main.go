package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contatInfo
}
type contatInfo struct {
	email string
	zip   int
}

func main() {
	contact := contatInfo{"rodolfo@gmail.com", 12345}
	rodolfo := person{"Rodolfo", "Leal", contact}

	alex := person{firstName: "Alex",
		lastName: "Anderson",
		contact:  contatInfo{"alex@mail.com", 672536}}

	var joao person

	joao.firstName = "Joao"
	joao.lastName = "Silva"

	fmt.Println(rodolfo, alex, joao)
	fmt.Printf("%+v", alex)
}
