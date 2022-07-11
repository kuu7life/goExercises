package main

import (
	"fmt"
	"math"
)

type powersuply interface {
	powerOfTwoNumbers() int
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type power struct {
	num, pow float64
}

type circle struct {
	radius float64
}

func (p power) powerOfTwoNumbers() int {
	return int(math.Pow(p.num, p.pow))
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func mess(p powersuply) {
	fmt.Println("powersuply: ", p)
	fmt.Println("power: ", p.powerOfTwoNumbers())
}

func measure(g geometry) {
	fmt.Println("geometry: ", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	p := power{num: 2, pow: 4}

	fmt.Println("Powersuply interface")
	mess(p)
	fmt.Println("+++++")

	fmt.Println("Geometry interface")
	measure(r)
	measure(c)
}
