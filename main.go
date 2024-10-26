package main

import "fmt"

func main() {
	var myName string = "x9z0"

	fmt.Printf("Hello My name is %s \n", myName)
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "meow")
	for index, value := range animals {
		fmt.Printf("this is my index %d and this is my animal %s \n", index, value)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++

	}
}
