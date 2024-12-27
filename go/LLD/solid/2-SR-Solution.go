package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Shape interface {
	area() float64
}

type circleN struct {
	radius float64
}

func (c circleN) area() float64 {
	return 3.14 * c.radius * c.radius
}

type squareN struct {
	length float64
}

func (s squareN) area() float64 {
	return s.length * s.length
}

type OutPutter struct{}

func (o OutPutter) printArea(s Shape) string {
	return fmt.Sprintf("Area: %f\n", s.area())
}

func (o OutPutter) jsonArea(s Shape) string {
	res := struct {
		Area float64 `json:"area"`
	}{
		Area: s.area(),
	}
	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}
func main() {
	o := OutPutter{}
	c := circleN{radius: 5}
	o.printArea(c)
	s := squareN{length: 5}
	o.printArea(s)
}
