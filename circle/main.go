package main

import (
	"fmt"
	"math"
)

const (
	PI = 3.14
)

type Circle struct {
	Radius float64
}

func (c *Circle) Circumference() float64 {
	return 2 * PI * c.Radius
}

func (c *Circle) Area() float64 {
	return PI * math.Pow(c.Radius, 2)
}

func main() {
	myCircle := Circle{
		Radius: 3,
	}
	fmt.Println(myCircle.Circumference())
	fmt.Println(myCircle.Area())
}
