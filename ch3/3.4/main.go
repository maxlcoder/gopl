// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320            // canvas size in pixels
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.4        // pixels per z unit

)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange...+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "image/svg+xml")
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	if r.Form["width"] != nil {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}
	if r.Form["height"] != nil {
		width, _ = strconv.Atoi(r.Form["height"][0])
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isFinite(ax) || isFinite(ay) ||
				isFinite(bx) || isFinite(by) ||
				isFinite(cx) || isFinite(cy) ||
				isFinite(dx) || isFinite(dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,xy).
	sx := float64(width/2) + (x-y)*cos30*float64(xyscale)
	sy := float64(height/2) + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func isFinite(f float64) bool {
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return true
	}
	return false
}
