// lissajous 产生随机利萨茹图形的GIF动画
// 测试 go run lissajous.go > out.gif
package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"os"
	"io"
	"image/gif"
	"image"
	"math"
)

var palette4 = []color.Color{color.White, color.Black}

const (
	whiteIndex4 = 0 // 画板中的第一种颜色
	blackIndex4 = 1 // 画板中的下一种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil {
				log.Print(err)
			}
			cycles := r.FormValue("cycles")
			if cycles != "" {
				cycles, err := strconv.Atoi(r.FormValue("cycles"))
				if err != nil {
					log.Fatal(err)
				}
				if cycles > 0 {
					lissajous4(w, cycles)
				}
			} else {
				fmt.Println("参数错误")
			}

		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous4(os.Stdout, 0)
}

func lissajous4(out io.Writer, cycles int)  {
	const (
		// cycles = 5 // 完整的 x 振荡器变化个数
		res = 0.001 // 角度分辨率
		size = 100 // 图像画布分布 [-size..+size]
		nframes = 64 // 动画中的帧数
		delay = 8 // 以 10ms 为单位的帧间延迟
	)
	fmt.Println(cycles)
	if cycles == 0 {
		cycles = 5;
	}
	freq := rand.Float64() * 3.0 // y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette4)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex4)
		}
		phase += 0.1
		anim.Delay =  append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误

}