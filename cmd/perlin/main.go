// See LICENSE.txt for licensing information.

package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/drbig/perlin"
)

var (
	fWidth, fHeight int
	fAlpha, fBeta   float64
	fN              int
	fSeed           int64
	fScale          float64
	fSmooth         bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] path\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.IntVar(&fWidth, "w", 320, "image width")
	flag.IntVar(&fHeight, "h", 256, "image height")
	flag.Float64Var(&fAlpha, "a", 2.0, "alpha factor")
	flag.Float64Var(&fBeta, "b", 2.0, "beta factor")
	flag.IntVar(&fN, "n", 1, "octave factor")
	flag.Int64Var(&fSeed, "r", time.Now().Unix(), "random seed")
	flag.Float64Var(&fScale, "s", 1.0, "scaling factor")
	flag.BoolVar(&fSmooth, "c", false, "smooth (continuous) gradient")
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)
	sx := fScale / float64(fWidth)
	sy := fScale / float64(fHeight)

	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	defer f.Close()

	g := perlin.NewGenerator(fAlpha, fBeta, fN, fSeed)
	img := image.NewGray(image.Rect(0, 0, fWidth, fHeight))

	for y := 0; y < fHeight; y++ {
		for x := 0; x < fWidth; x++ {
			idx := (y-img.Rect.Min.Y)*img.Stride + (x-img.Rect.Min.X)*1
			nx := float64(x) * sx
			ny := float64(y) * sy
			if fSmooth {
				img.Pix[idx] = uint8(254.0 * ((g.Noise2D(nx, ny) + 1.0) / 2.0))
			} else {
				img.Pix[idx] = uint8(254.0 * g.Noise2D(nx, ny))
			}
		}
	}

	err = png.Encode(f, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
