package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) changeName(newName string) {
	p.Name = newName
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		age:  age,
	}
}

func main() {
	myPerson := Person{
		Name: "x9z0",
		age:  21,
	}
	myPerson.changeName("meow")

	fmt.Printf("this is my person %+v", myPerson)

}
