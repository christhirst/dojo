package main

import "fmt"

type person struct {
	vorname  string
	nachname string
}

func (p person) GetVorname() string {
	return p.vorname
}

func create(vorname, nachname string) *person {
	p := &person{
		vorname,
		nachname,
	}
	return p
}

func main() {
	p := create("myname", "nachname")
	fmt.Println(p.GetVorname())
}
