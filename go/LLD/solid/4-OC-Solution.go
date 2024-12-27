package main

/*
Open-Closed Principle (OCP):
The Open-Closed Principle states that software entities (classes, modules, functions) should be open for extension but closed for modification.
This means that we should design our code in a way that allows us to add new functionality without modifying existing code.
By using techniques such as inheritance, interfaces, and dependency injection, we can achieve this principle and avoid introducing bugs or breaking existing code when making changes.
*/
type Shapes interface {
	Area() float64
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func (t triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func AreaSum(shapes []Shapes) float64 {
	area := 0.0
	for _, shape := range shapes {
		area += shape.Area()
	}
	return area
}
