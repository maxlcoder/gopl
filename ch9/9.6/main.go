package main

import (
	"github.com/maxlcoder/gopl/ch9/9.6/mandelbrot"
	"image/png"
	"log"
	"os"
	"runtime"
)

func main() {
	workers := runtime.GOMAXPROCS(-1)
	img := mandelbrot.ConcurrentRender(workers)

	f, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, img)
}