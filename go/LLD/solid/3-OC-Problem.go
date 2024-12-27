package main

import "fmt"

/*
Open-Closed Principle (OCP):
The Open-Closed Principle states that software entities (classes, modules, functions) should be open for extension but closed for modification.
This means that we should design our code in a way that allows us to add new functionality without modifying existing code.
By using techniques such as inheritance, interfaces, and dependency injection, we can achieve this principle and avoid introducing bugs or breaking existing code when making changes.
*/
type rectangle struct {
	length float64
	width  float64
}

type triangle struct {
	base   float64
	height float64
}

/*
Open-Closed Principle (OCP):
The code violates this principle. If you wanted to add a new shape, you would need to modify the `sumArea` function to add a new case in the switch statement.
A better approach would be to define an interface (e.g., `Shape`) that includes an `Area` method, and have `rectangle` and `triangle` implement this interface.
This way, you could add new shapes without modifying the `sumArea` function.
*/
func sumArea(shapes []any) float64 {
	area := 0.0
	for _, shape := range shapes {
		switch shape.(type) {
		case rectangle:
			area += shape.(rectangle).length * shape.(rectangle).width
		case triangle:
			area += 0.5 * shape.(triangle).base * shape.(triangle).height
		}
	}
	return area
}

func main() {
	r := rectangle{length: 5, width: 3}
	t := triangle{base: 4, height: 2}
	shapes := []any{r, t}
	fmt.Println(sumArea(shapes))
}
