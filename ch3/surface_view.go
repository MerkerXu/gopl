package main

import (
    "net/http"
    "io"
	"math"
    "log"
    "fmt"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")
        surface(w)
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("10.200.247.243:8000", nil))
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, abool := corner(i+1, j)
			bx, by, bbool := corner(i, j)
			cx, cy, cbool := corner(i, j+1)
			dx, dy, dbool := corner(i+1, j+1)
            if abool && bbool && cbool && dbool {
                fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                    ax, ay, bx, by, cx, cy, dx, dy)
            }
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//compute surface height z
	z := f(x, y)
    if math.IsNaN(z) {
        return 0.,0.,false
    }

	// projet (x, y, z) isometrically onto 2-D SVG canvas(sx, sv)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
