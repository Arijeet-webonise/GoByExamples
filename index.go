package main

import (
	"fmt"
	"math"
)

type soild interface {
	surfacearea() float64
	volume() float64
}

type cube struct {
	side float64
}

type cuboid struct {
	length float64
	width  float64
	height float64
}

type sphere struct {
	radius float64
}

func (c cube) volume() float64 {
	return c.side * c.side * c.side
}

func (c cuboid) volume() float64 {
	return c.width * c.length * c.height
}

func (s sphere) volume() float64 {
	return (4 / 3) * math.Pi * s.radius * s.radius * s.radius
}

func (c cube) surfacearea() float64 {
	return 6 * c.side * c.side
}

func (c cuboid) surfacearea() float64 {
	return 2 * (c.length*c.height + c.length*c.width + c.height*c.width)
}

func (s sphere) surfacearea() float64 {
	return 4 * math.Pi * s.radius * s.radius
}

func measure(s soild) {
	fmt.Println(s)
	fmt.Println(s.surfacearea())
	fmt.Println(s.volume())
}

func main() {
	s := sphere{radius: 5}
	c1 := cube{side: 5}
	c2 := cuboid{length: 5, width: 2, height: 3}

	measure(s)
	measure(c1)
	measure(c2)
}
