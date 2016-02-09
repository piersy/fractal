package main

import (
	"image"
	"testing"
)

// from fib_test.go
func BenchmarkFractalGeneration(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		im := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{1000, 1000}})
		GenerateFractal(im)
	}
}
