package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

type RequestParam struct {
	height, width float64
	color         string
}

type SurfaceParts struct {
	width, height, xyrange, xyscale, zscale, angle, sin30, cos30 float64
	cells                                                        int
	color                                                        string
}

func NewSurfaceParts(w, h float64) *SurfaceParts {
	s := SurfaceParts{
		width:   w,
		height:  h,
		cells:   100,
		xyrange: 30.0,
		zscale:  h * 0.4,
		angle:   math.Pi / 6,
	}
	s.xyscale = w / 2 / s.xyrange
	s.sin30 = math.Sin(s.angle)
	s.cos30 = math.Cos(s.angle)

	return &s
}

func main() {
	http.HandleFunc("/surface", surfaceHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")
	q := r.URL.Query()
	wid := float64(600)
	if q["width"] != nil {
		wid, _ = strconv.ParseFloat(q["width"][0], 64)
	}
	hei := float64(320)
	if q["height"] != nil {
		hei, _ = strconv.ParseFloat(q["height"][0], 64)
	}
	col := "ffffff"
	if q["color"] != nil {
		col = q["color"][0]
	}
	// fmt.Println("ここ注目")
	// fmt.Println(q["width"][0], q["height"][0])
	p := RequestParam{
		width:  wid,
		height: hei,
		color:  col,
	}
	surface(w, p)

}

func surface(out io.Writer, p RequestParam) {

	s := NewSurfaceParts(p.width, p.height)
	s.color = p.color

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", s.width, s.height)

	for i := 0; i < s.cells; i++ {
		for j := 0; j < s.cells; j++ {
			ax, ay := corner(i+1, j, s)
			bx, by := corner(i, j, s)
			cx, cy := corner(i, j+1, s)
			dx, dy := corner(i+1, j+1, s)
			fmt.Fprintf(out, "<polygon points='%g, %g, %g, %g, %g, %g, %g, %g' style='fill:#%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, s.color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, s *SurfaceParts) (float64, float64) {
	// マス目(i, j)の角のてん(x, y)を見つける
	x := s.xyrange * (float64(i)/float64(s.cells) - 0.5)
	y := s.xyrange * (float64(j)/float64(s.cells) - 0.5)

	// 面の高さzを計算する
	z := f(x, y)

	// (x, y, z)を2-D SVGキャンバス(sx, sy)へ等角的に東映
	sx := s.width/2 + (x-y)*s.cos30*s.xyscale
	sy := s.height/2 + (x+y)*s.sin30*s.xyscale - z*s.zscale
	return sx, sy
}

// func f(x, y float64) float64 {
// 	r := math.Hypot(x, y) // (0, 0)からの距離
// 	return math.Sin(r) / r
// }

func f(x, y float64) float64 {
	// return math.Sin((x * y / 10) / 10)
	return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
}
