package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0, 153, 51, 1}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {

	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	lissajous(w)
	// }
	http.HandleFunc("/", lessajousHandler) // 個々のリクエストに対してhandlerが呼ばれる
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

func lessajousHandler(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query()

	fmt.Println(r.URL.Query, k["cycles"])
	lissajous(w, k["cycles"][0])
}

func lissajous(out io.Writer, s string) {
	var (
		cycles  int
		res     float64
		size    int
		nframes int
		delay   int
	)
	cycles = 5   // number of complete x oscillator revolutions
	res = 0.001  // angular resolution
	size = 100   // image canvas covers [-size..+size]
	nframes = 64 // number of animation frames
	delay = 8    // delay between frames in 10ms units

	if s != "" {
		cycles, _ = strconv.Atoi(s)
	}

	// freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	freq := 0.90
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < (float64)(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*(float64)(size)+0.5), size+int(y*(float64)(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
