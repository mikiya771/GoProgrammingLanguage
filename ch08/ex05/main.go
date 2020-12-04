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
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := make(chan struct{})
	for py := 0; py < height; py++ {
		go func(py int) {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbort(z))
			}
			wg <- struct{}{}
		}(py)
	}
	for py := 0; py < height; py++ {
		<-wg
	}
	png.Encode(os.Stdout, img)
}
func mandelbort(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			var red uint8 = IntMax(255-contrast*n, 0)
			var green uint8 = IntMax(255-contrast*IntMax(0, n-5), 0)
			var blue uint8 = IntMax(255-contrast*IntMax(0, n-10), 0)
			return color.RGBA{red, green, blue, 255}
		}
	}
	return color.Black
}
func IntMax(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
