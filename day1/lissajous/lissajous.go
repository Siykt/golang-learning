package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 用时间戳设置随机数
	rand.Seed(time.Now().UTC().UnixNano())
	// 标准输出 (保存文件)
	// lissajous(os.Stdout)
	// 生成文件
	file, err := os.Create("lissajous.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lissajous(file)
}

func lissajous(out io.Writer) {
	var palette = []color.Color{color.RGBA{2, 197, 51, 1}, color.White}

	// 使用 const 声明常量
	const (
		whiteIndex = 0      // first color in palette
		blackIndex = 1      // next color in palette
		cycles     = 5      // number of complete x oscillator revolutions
		res        = 0.0001 // angular resolution
		size       = 100    // image canvas covers [-size..+size]
		nframes    = 64     // number of animation frames
		delay      = 8      // delay between frames in 10ms units
	)

	// 随机的振率
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	// GIF图像
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// 生成每一帧的图片
	for i := 0; i < nframes; i++ {
		// 生成矩形图片
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 填色
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 设置像素点颜色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1

		// append 动画间隔(GIF时间轴)
		anim.Delay = append(anim.Delay, delay)
		// append 当前帧图片
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
