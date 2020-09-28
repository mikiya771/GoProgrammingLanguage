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
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbort(z))
		}
	}
	png.Encode(os.Stdout, SuperSampling(img))
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

func SuperSampling(img *image.RGBA) *image.RGBA {
	height := img.Rect.Max.Y*2 - 1
	width := img.Rect.Max.X*2 - 1
	newimg := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			sampleA := img.RGBAAt(px/2, py/2)
			sampleB := img.RGBAAt((px+1)/2, py/2)
			sampleC := img.RGBAAt(px/2, (py+1)/2)
			sampleD := img.RGBAAt((px+1)/2, (py+1)/2)
			R := sampleA.R/4 + sampleB.R/4 + sampleC.R/4 + sampleD.R/4
			G := sampleA.G/4 + sampleB.G/4 + sampleC.G/4 + sampleD.G/4
			B := sampleA.B/4 + sampleB.B/4 + sampleC.B/4 + sampleD.B/4
			newimg.Set(px, py, color.RGBA{R, G, B, 255})
		}
	}
	return newimg
}