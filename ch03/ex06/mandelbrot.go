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
		width, height          = 1024, 1024
	)
	bimg := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点(px, py)は複素数値zを表している
			bimg.SetRGBA(px, py, mandelbrot(z))
		}
	}
	aimg := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			c1 := bimg.RGBAAt(px, py)
			c2 := bimg.RGBAAt(px+1, py)
			c3 := bimg.RGBAAt(px, py+1)
			c4 := bimg.RGBAAt(px+1, py+1)

			ravg := (float32(c1.R) + float32(c2.R) + float32(c3.R) + float32(c4.R)) / 4
			gavg := (float32(c1.G) + float32(c2.G) + float32(c3.G) + float32(c4.G)) / 4
			bavg := (float32(c1.B) + float32(c2.B) + float32(c3.B) + float32(c4.B)) / 4
			aavg := (float32(c1.A) + float32(c2.A) + float32(c3.A) + float32(c4.A)) / 4
			aimg.SetRGBA(px, py, color.RGBA{uint8(ravg), uint8(gavg), uint8(bavg), uint8(aavg)})
		}
	}
	png.Encode(os.Stdout, aimg) // TODO エラー処理
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - contrast*n
			return color.RGBA{r, 0, 0, 255}
			// return color.Gray{255 - contrast*n}
		}
	}
	// return color.Black
	return color.RGBA{0, 255, 0, 255}
}
