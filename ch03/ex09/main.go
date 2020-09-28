package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", newtonPic)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func newtonPic(w http.ResponseWriter, r *http.Request) {
	queryMap := map[string]float64{}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if i, err := strconv.ParseFloat(v[0], 64); err != nil {
			fmt.Print(err)
		} else {
			queryMap[k] = i
		}
	}
	fmt.Println(queryMap)
	if queryMap["xrange"] == 0 {
		queryMap["xrange"] = 2
	}
	if queryMap["yrange"] == 0 {
		queryMap["yrange"] = 2
	}
	if queryMap["percent"] == 0 {
		queryMap["percent"] = 100
	}
	imgWriter(w, newton, queryMap["xrange"], queryMap["yrange"], queryMap["percent"])
}
func imgWriter(w http.ResponseWriter, f func(complex128) color.Color, xrange float64, yrange float64, percent float64) {
	xmin, ymin, xmax, ymax := -xrange, -yrange, xrange, yrange
	width, height := int(1024*percent/100), int(1024*percent/100)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
	log.Println(xrange, yrange, percent)
}

func newton(z complex128) color.Color {
	const iter = 15
	const contr = 16
	for i := uint8(0); i < iter; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(1-z) < 1e-6 {
			return color.Gray{255 - contr*i}
		}
		if cmplx.Abs(-1-z) < 1e-6 {
			return color.Gray{255 - contr*i}
		}
		if cmplx.Abs(1i-z) < 1e-6 {
			return color.Gray{255 - contr*i}
		}
		if cmplx.Abs(-1i-z) < 1e-6 {
			return color.Gray{255 - contr*i}
		}
	}
	return color.Black
}
