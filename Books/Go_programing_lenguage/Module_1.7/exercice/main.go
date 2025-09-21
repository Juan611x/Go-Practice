package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, // blanco
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // rojo
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // verde
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // azul
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // amarillo
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // cian
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // magenta
	color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, // naranja
	color.RGBA{0x8A, 0x2B, 0xE2, 0xFF}, // violeta
	color.RGBA{0xFF, 0x69, 0xB4, 0xFF}, // rosa
	color.RGBA{0x80, 0x80, 0x80, 0xFF}, // gris
	color.RGBA{0x00, 0x00, 0x00, 0xFF}, // negro
}

const (
	whiteIndex = 0 //first color in palette
	blackIndex = 1
)

// go run .\Books\Go_programing_lenguage\Module_1.7\exercice\main.go
//URL: http://localhost:8000?cycles=20

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := 5.0
		if val := r.URL.Query().Get("cycles"); val != "" {
			if c, err := strconv.ParseFloat(val, 64); err == nil {
				cycles = c
			}
		}
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func lissajous(out io.Writer, cycles float64) {
	const (
		// cycles = 5     //number of complete x oscillator revolutions
		res    = 0.001 //angular resolution
		size   = 100   //image canvas covers [-size..+size]
		nframe = 64    //number of anumation frames
		delay  = 8     //delay between frames in 10ms unit
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscilator
	anim := gif.GIF{LoopCount: nframe}
	phase := 0.0 //phase difference
	for i := 0; i < nframe; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Int()%len(palette)+1))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
