package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Path 是连接多个点的直线段
type Path []Point


// 普通的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point类型的方法
func (p Point) Distance(q Point) float64  {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance 方法返回路径的长度
func (path Path) Distance() float64  {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1,  2}
	q := Point{4, 8}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	r := &Point{1, 2}
	r.ScaleBy(3)
	fmt.Println(r)
	fmt.Println(*r)

	r1 := Point{1, 2}
	r1.ScaleBy(3)
	fmt.Println(r1)

	t := Point{1, 5}
	(&t).ScaleBy(2)
	fmt.Println(t)

	s := Point{2, 5}
	s.ScaleBy(2)
	fmt.Println(s)

}