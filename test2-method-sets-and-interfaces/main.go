package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area :", s.area())
}

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	c := circle{
		radius: 8,
	}
	info(&c)

	p1 := person{"Logan"}
	saySomething(&p1)
	p1.speak()
}
