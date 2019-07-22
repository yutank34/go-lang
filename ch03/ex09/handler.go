package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/newton", newtonHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func newtonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	xmin, ymin, xmax, ymax := -1, -1, +1, +1
	width, height := 1024, 1024

	q := r.URL.Query()
	if q["magnification"] != nil {
		m, _ := strconv.Atoi(q["magnification"][0])
		xmin, ymin, xmax, ymax = -m, -m, +m, +m
	}

	if q["width"] != nil {
		w, err := strconv.Atoi(q["width"][0])
		if err != nil {
			log.Fatal(err)
		}
		width = w
	}

	if q["height"] != nil {
		h, err := strconv.Atoi(q["height"][0])
		if err != nil {
			log.Fatal(err)
		}
		height = h
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(w, img)
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
