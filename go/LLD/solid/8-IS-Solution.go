package main

/*
The Interface Segregation Principle states that clients should not be forced to depend on interfaces they do not use.
It promotes the idea of creating smaller, more focused interfaces rather than large, monolithic ones.
By doing so, we can avoid unnecessary dependencies and make our code more modular and maintainable.
*/

type ShapesWithArea interface {
	area() float64
}
type ShapesWithVolume interface {
	volume() float64
}

type CirclE struct {
	radius float64
}

func (c CirclE) area() float64 {
	return 3.14 * c.radius * c.radius
}

type CubE struct {
	side float64
}

func (cu CubE) volume() float64 {
	return cu.side * cu.side * cu.side
}

func totalVolume(s []ShapesWithVolume) float64 {
	volume := 0.0
	for _, shape := range s {
		volume += shape.volume()
	}
	return volume
}
func totalArea(s []ShapesWithArea) float64 {
	area := 0.0
	for _, shape := range s {
		area += shape.area()
	}
	return area
}
