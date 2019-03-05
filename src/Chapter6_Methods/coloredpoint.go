package main

import (
	"fmt"
	"image/color"
	"math"
)
type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64  {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point)Distance(q Point) float64  {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }
type Path []Point
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main()  {
	var cp ColorPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255,0,0, 255}
	blue := color.RGBA{0,0, 255, 255}
	var p = ColorPoint{Point{1,1,}, red}
	var q = ColorPoint{Point{5,4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	p1 := Point{1,2}
	q1 := Point{4, 6}

	distanceFromP := p.Distance;
	fmt.Println(distanceFromP(q1))
	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p1.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)


}
