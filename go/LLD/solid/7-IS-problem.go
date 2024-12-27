package main

/*
The Interface Segregation Principle states that clients should not be forced to depend on interfaces they do not use.
It promotes the idea of creating smaller, more focused interfaces rather than large, monolithic ones.
By doing so, we can avoid unnecessary dependencies and make our code more modular and maintainable.
*/

type NewShape interface {
	area() float64
	volume() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

/*
Interface Segregation Principle (ISP): The code violates this principle. 
The `NewShape` interface includes both `area` and `volume` methods, but not all shapes need both. 
For example, the `Circle` struct has to implement a `volume` method that returns 0, because a circle doesn't have a volume. 
It would be better to split the `NewShape` interface into two interfaces, one for shapes that have an area and another for shapes that have a volume.
*/
func (c Circle) volume() float64 {
	return 0
}

type Cube struct {
	side float64
}

func (cu Cube) area() float64 {
	return 0
}
func (cu Cube) volume() float64 {
	return cu.side * cu.side * cu.side
}

func TotalVolume(s []NewShape) float64 {
	volume := 0.0
	for _, shape := range s {
		volume += shape.volume()
	}
	return volume
}
func TotalArea(s []NewShape) float64 {
	area := 0.0
	for _, shape := range s {
		area += shape.area()
	}
	return area
}
