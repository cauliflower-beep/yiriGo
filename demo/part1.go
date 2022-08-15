package main

var p = newPerson()

func newPerson() *person {
	return &person{}
}

type person struct{}
