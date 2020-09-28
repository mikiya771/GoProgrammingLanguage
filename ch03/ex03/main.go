package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 600
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.1
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			a, ax, ay := corner(i+1, j)
			b, bx, by := corner(i, j)
			c, cx, cy := corner(i, j+1)
			d, dx, dy := corner(i+1, j+1)
			color := "white"
			if calcDot(i, j)*calcDot(i+1, j+1) <= 0 {
				x1 := xyrange * (float64(i)/cells - 0.5)
				y1 := xyrange * (float64(j)/cells - 0.5)
				x2 := xyrange * (float64(i+1)/cells - 0.5)
				y2 := xyrange * (float64(j+1)/cells - 0.5)
				if math.Hypot(x1, y1) < math.Hypot(x2, y2) {
					if calcDot(i, j) > 0 {
						color = "red"
					} else {
						color = "blue"
					}
				} else {
					if calcDot(i, j) > 0 {
						color = "blue"
					} else {
						color = "red"
					}
				}
			}
			if a && b && c && d {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s;fill-rule:nonzero;'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}
func corner(i, j int) (bool, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	s := true
	if math.IsNaN(z) || math.IsInf(z, 0) {
		s = false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*cos30*xyscale - z*zscale
	return s, sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func calcDot(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	r := math.Hypot(x, y)
	return math.Sin(r) + r*math.Cos(r)
}
