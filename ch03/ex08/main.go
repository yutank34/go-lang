package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	base                   = 0.5
	xmin, ymin, xmax, ymax = -base, -base, +base, +base
	width, height          = 1024, 1024
)

func main() {
	img := fractal3()
	png.Encode(os.Stdout, img)
}

func fractal1() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot1(complex64(z)))
		}
	}
	return img
}

func mandelbrot1(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func fractal2() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot2(z))
		}
	}
	return img
}

func mandelbrot2(z complex128) color.Color {
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

type FloatComplex struct {
	real, imag *big.Float
}

func (z *FloatComplex) Add(x, y *FloatComplex) *FloatComplex {
	z.real = new(big.Float).Add(x.real, y.real)
	z.imag = new(big.Float).Add(x.imag, y.imag)
	return z
}

func (z *FloatComplex) Mul(x, y *FloatComplex) *FloatComplex {
	z.real = new(big.Float).Sub(
		new(big.Float).Mul(x.real, y.real),
		new(big.Float).Mul(x.imag, y.imag))
	z.imag = new(big.Float).Add(
		new(big.Float).Mul(x.real, y.imag),
		new(big.Float).Mul(x.imag, y.real))
	return z
}

func (z *FloatComplex) Abs(x *FloatComplex) *big.Float {
	return new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(x.real, x.real),
			new(big.Float).Mul(x.imag, x.imag),
		),
	)
}

func fractal3() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := &FloatComplex{big.NewFloat(x), big.NewFloat(y)}
			img.Set(px, py, mandelbrot3(z))
		}
	}
	return img
}

func mandelbrot3(z *FloatComplex) color.Color {
	const iterations = 200
	const contrast = 15

	v := &FloatComplex{big.NewFloat(0), big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		v = new(FloatComplex).Add(
			new(FloatComplex).Mul(v, v),
			z,
		)

		if new(FloatComplex).Abs(v).Cmp(big.NewFloat(2)) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

type RatComplex struct {
	real, imag *big.Rat
}

func (z *RatComplex) Add(x, y *RatComplex) *RatComplex {
	z.real = new(big.Rat).Add(x.real, y.real)
	z.imag = new(big.Rat).Add(x.imag, y.imag)
	return z
}

func (z *RatComplex) Mul(x, y *RatComplex) *RatComplex {
	z.real = new(big.Rat).Sub(
		new(big.Rat).Mul(x.real, y.real),
		new(big.Rat).Mul(x.imag, y.imag))
	z.imag = new(big.Rat).Add(
		new(big.Rat).Mul(x.real, y.imag),
		new(big.Rat).Mul(x.imag, y.real))
	return z
}

// func (z *RatComplex) Abs(x *RatComplex) *big.Rat {
// 	return new(big.Rat).Sqrt(
// 		new(big.Rat).Add(
// 			new(big.Rat).Mul(x.real, x.real),
// 			new(big.Rat).Mul(x.imag, x.imag),
// 		),
// 	)
// }
