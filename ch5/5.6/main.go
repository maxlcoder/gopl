//  surface 函数根据一个三维的曲面函数计算并生成 SVG
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 // 以像素表示曲面画布
	cells = 100 // 网格单元的个数
	xyrange = 30.0 // 坐标轴的范围 (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange // x 或 y 轴上每个单位长度的像素
	zscale = height * 0.4 // z 轴上每个单位长度的像素
	angle = math.Pi / 6 // x，y轴的角度（=30度）
)

var sin30, cos30 =  math.Sin(angle), math.Cos(angle) // sim(30度)，cos(30度)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'> "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j +1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy float64) {
	// 求出网格单元（i,j）的顶点坐标（x,y）
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算曲面高估 z
	z := f(x, y)

	// 将 (x,y,z) 等角投射到二维 SVG 绘图平面上， 坐标是 （sx, sy）
	sx = width/2 + (x-y)*cos30*xyrange
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}