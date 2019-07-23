// lissajous 产生随机利萨茹图形的GIF动画
// 测试 go run lissajous.go > out.gif
package main

import (
	"image/color"
	"log"
	"math/rand"
	"net/http"
	"time"
	"os"
	"io"
	"image/gif"
	"image"
	"math"
)

// 定义3种颜色随机使用
var palette3 = []color.Color{color.RGBA{0xFF, 0xDE, 0xAD, 0xFF}, color.RGBA{0xEE, 0x5C, 0x42, 0xFF}, color.RGBA{0xB3, 0xEE, 0x3A, 0xFF}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous3(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous3(os.Stdout)
}

func lissajous3(out io.Writer)  {
	const (
		cycles = 5 // 完整的 x 振荡器变化个数
		res = 0.001 // 角度分辨率
		size = 100 // 图像画布分布 [-size..+size]
		nframes = 64 // 动画中的帧数
		delay = 8 // 以 10ms 为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette3)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(3)))
		}
		phase += 0.1
		anim.Delay =  append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误

}