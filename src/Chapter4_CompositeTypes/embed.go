package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Whell struct {
	Circle
	Spokes int
}

func main()  {
	var w Whell;
	w.X = 8
	w.X = 8
	w.Radius = 5
	w.Spokes = 20

	w = Whell{Circle{Point{8, 8}, 5}, 20}

	w = Whell{
		Circle: Circle{
			Point: Point{X:8, Y:8},
			Radius:5,
		},
		Spokes:20,
	}

	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)

}
