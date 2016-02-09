package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
)

var (
	centerx = flag.Float64("centerx", 0, "X component of center")
	centery = flag.Float64("centery", 0, "y component of center")
	width   = flag.Float64("width", 4, "viewport width")
	height  = flag.Float64("height", 4, "viewport height")
)

type cn struct {
	i float64
	r float64
}

func main() {

	flag.Parse()
	im := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{1000, 1000}})
	z := &cn{}

	c := &cn{}
	g := color.Gray{}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {

			c.r = (*width * float64(x) / float64(1000)) - *width/float64(2)
			c.i = (*height * float64(y) / float64(1000)) - *height/float64(2)
			c.r += *centerx
			c.i += *centery

			z.r = 0
			z.i = 0
			zrsqr := z.r * z.r
			zisqr := z.i * z.i

			var i uint8 = 0
			for zrsqr+zisqr < 4.0 && i < 255 {
				i++
				z.i = square(z.r+z.i) - zrsqr - zisqr
				z.i += c.i
				z.r = zrsqr - zisqr + c.r
				zrsqr = square(z.r)
				zisqr = square(z.i)
			}
			g.Y = i
			im.SetGray(x, y, g)

		}
	}

	f, err := os.Create("./output.png")
	defer f.Close()
	if err != nil {

		println(err)
	}
	png.Encode(f, im)
}
func square(x float64) float64 {
	return x * x
}
