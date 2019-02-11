package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contatInfo
}
type contatInfo struct {
	email string
	zip   int
}

func main() {
	contact := contatInfo{"rodolfo@gmail.com", 12345}
	rodolfo := person{"Rodolfo", "Leal", contact}

	alex := person{firstName: "Alex",
		lastName:   "Anderson",
		contatInfo: contatInfo{"alex@mail.com", 672536}}

	var joao person

	joao.firstName = "Joao"
	joao.lastName = "Silva"
	joao.contatInfo = contatInfo{email: "jao@mail.com", zip: 763782}

	rodolfo.print()
	alex.print()
	joao.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
