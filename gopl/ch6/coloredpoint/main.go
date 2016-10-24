package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(s float64) {
	p.X *= s
	p.Y *= s
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Y = 2
	fmt.Println(cp.Point.X)
	fmt.Println(cp.Point.Y)

	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance
	fmt.Println(distance(p, q))

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}
