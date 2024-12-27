package main

import "fmt"

/*
Single Responsibility Principle (SRP):
This principle states that a class or module should have only one reason to change.
In other words, a class should have a single responsibility or purpose.
By adhering to this principle, we can ensure that each class is focused on a specific task, making it easier to understand, test, and maintain.
*/

type circle struct {
	radius float64
}

// Problem here is that it is printing and calculating area
func (c *circle) area() {
	fmt.Printf("Circle area %f\n", 3.14*c.radius*c.radius)
}

type square struct {
	length float64
}

func (s *square) area() {
	fmt.Printf("Circle area %f\n", s.length*s.length)
}

func main() {
	c := circle{radius: 5}
	c.area()
	s := square{length: 5}
	s.area()
}
