package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var width, height = 600, 600
var xyscale = float64(width) / 2 / xyrange
var zscale = float64(height) * 0.1

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	formMap := map[string]string{}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		formMap[k] = v[0]
	}
	color := string(getMap(formMap, "color", "white"))
	width, _ = strconv.Atoi(getMap(formMap, "width", "600"))
	height, _ = strconv.Atoi(getMap(formMap, "height", "600"))
	xyscale = float64(width) / 2 / xyrange
	zscale = float64(height) * 0.1
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: %s; stroke-width: 0.7' "+"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			a, ax, ay := corner(i+1, j)
			b, bx, by := corner(i, j)
			c, cx, cy := corner(i, j+1)
			d, dx, dy := corner(i+1, j+1)
			if a && b && c && d {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (bool, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	s := true
	if math.IsNaN(z) || math.IsInf(z, 0) {
		s = false
	}

	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*cos30*xyscale - z*zscale
	return s, sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
func getMap(m map[string]string, key string, def string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return def
}
