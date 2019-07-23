package main

import "fmt"

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

type Circle struct {
	// X, Y, Radius int
	// Center Point
	Point
	Radius int
}

type Wheel struct {
	// X, Y, Radius, Spokes int
	// Circle Circle
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)
	// 输出
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // 注意，尾部的逗号是必需的（Radius后面的都要也一样）
	}
	fmt.Printf("%#v\n", w)

	w.X = 42

	fmt.Printf("%#v\n", w)
}
