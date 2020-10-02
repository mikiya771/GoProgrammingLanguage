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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}
func newton(z complex128) color.Color {
	const iter = 15
	const contr = 16
	for i := uint8(0); i < iter; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(1-z) < 1e-6 {
			return color.RGBA{0, 0, 255 - contr*i, 255}
		}
		if cmplx.Abs(-1-z) < 1e-6 {
			return color.RGBA{0, 255 - contr*i, 0, 255}
		}
		if cmplx.Abs(1i-z) < 1e-6 {
			return color.RGBA{255 - contr*i, 0, 0, 255}
		}
		if cmplx.Abs(-1i-z) < 1e-6 {
			return color.RGBA{0, 255 - contr*i, 255 - contr*i, 255}
		}
	}
	return color.Black
}
