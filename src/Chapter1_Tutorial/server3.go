package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var palette = []color.Color{color.White, color.RGBA{0xff, 0x00, 0x00,0xff},
	color.RGBA{0x00, 0xff, 0x00,0xff}, color.RGBA{0x00, 0x00, 0xff,0xff}}

const(
	whiteIndex = 0 //first color in palette
	RedIndex = 1 //next color in palette
	GreenIndex = 2
	BlueInde = 3
)

func main()  {

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil{
			log.Print(err)
		}
		for k ,v := range r.Form{
			if k == "cycles"{
				cycles, err := strconv.Atoi(v[0])
				if err != nil{
					fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				}
				Lissajous(w, cycles)
			}
		}

	}
	http.HandleFunc("/", handler2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}




func handler2(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil{
		log.Print(err)
	}
	for k ,v := range r.Form{
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func Lissajous(out io.Writer, cycle int )  {

	cycles	:= cycle		//number of complete x oscillator revolutions
	const(
		res		= 0.001 // angular resolution
		size	= 100   // image canvas covers [-size..+size]
		nframes	= 64	// number of animation frames
		delay	= 8 		// delay between frames in 10ms units
	)
	freq := rand.Float64()*3.0
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0
	for i := 0; i < nframes; i++{
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t:=0.0; t<float64(cycles)*2*math.Pi; t+=res{
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), BlueInde)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
