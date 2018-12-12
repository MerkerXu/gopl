package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "net/http"
    "io"
    "log"
)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        mandelbrotPng(w)
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("10.200.247.243:8000", nil))
}

func mandelbrotPng(out io.Writer) {
    const (
        xmin, ymin, xmax, ymax  = -2, -2, +2, +2
        width, height           = 1024, 1024
    )
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py :=0; py<height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px:=0; px<width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            //Image point (px,py) represents complex value z
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(out, img) //NOTE: ignoring errors
}

func mandelbrot(z complex128) color.RGBA {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n:=uint8(0); n<iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return color.RGBA{0xFF - contrast*n, 0xFF - contrast*n, 0xFF, 0xFF}
        }
    }
    return color.RGBA{0,0,0,0xFF}
}
