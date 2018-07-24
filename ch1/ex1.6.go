package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
)

var palette = []color.Color{color.White, color.RGBA{255,0,0,255},
                    color.RGBA{0,255,0,255}, color.RGBA{0,0,255,255},
                    color.RGBA{255,255,255,255}}

const (
    whiteIndex = 0
    redIndex = 1
    greenIndex = 2
    blueIndex = 3
    blackIndex = 4
)

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5
        res     = 0.001
        size    = 100
        nframes = 64
        delay   = 8
    )

    freq := rand.Float64()*3.0
    anim := gif.GIF{LoopCount: nframes}
    phase:= 0.0

    for i:=0; i<nframes; i++ {
        rect := image.Rect(0,0,2*size+1,2*size+1)
        img  := image.NewPaletted(rect, palette)
        for t:= 0.0; t<cycles*2*math.Pi; t+=res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            colorIndex := getColorIndex()
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                colorIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}

func getColorIndex() (uint8) {
    indexRand := rand.Float64() * 4.0
    if indexRand < 1.0 {
        return redIndex
    }
    if indexRand < 2.0 {
        return greenIndex
    }
    if indexRand < 3.0 {
        return blueIndex
    }
    return blackIndex
}
