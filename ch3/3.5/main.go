package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, acos(z))
		}
	}
	png.Encode(os.Stdout, img)

}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v) * 128) + 127
	red := uint8(imag(v) * 128) + 127
	return color.YCbCr{192,blue,red}
}