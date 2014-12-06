// See LICENSE.txt for licensing information.

// Package perlin implements reusable Perlin noise generators.
package perlin

import (
	"math"
	"math/rand"
)

// Magic numbers, not commented even two steps upstream.
const (
	cb  = 0x100
	cn  = 0x1000
	cbm = 0xff
)

// Generator holds Perlin noise parameters and random seed.
type Generator struct {
	Alpha float64 // Alpha is the weight factor used during summing
	Beta  float64 // Beta is the harmonic scaling/spacing factor
	N     int     // N is the number of octaves/samples
	Seed  int64   // Seed is the number used to seed the RNG

	r *rand.Rand
	p [cb + cb + 2]int

	g3 [cb + cb + 2][3]float64
	g2 [cb + cb + 2][2]float64
	g1 [cb + cb + 2]float64
}

// NewGenerator returns seeded Generator for given parameters.
func NewGenerator(alpha, beta float64, n int, seed int64) *Generator {
	g := &Generator{
		Alpha: alpha,
		Beta:  beta,
		N:     n,
	}
	g.Reset(seed)

	return g
}

// Reset re-seeds an existing Generator.
// Note that you can change other parameters before calling this, effectively completely re-parametrising the Generator.
func (g *Generator) Reset(seed int64) {
	g.Seed = seed
	g.r = rand.New(rand.NewSource(seed))

	for i := 0; i < cb; i++ {
		g.p[i] = i
		g.g1[i] = float64((g.r.Int()%(cb+cb))-cb) / cb

		for j := 0; j < 2; j++ {
			g.g2[i][j] = float64((g.r.Int()%(cb+cb))-cb) / cb
		}
		normalize2(&g.g2[i])

		for j := 0; j < 3; j++ {
			g.g3[i][j] = float64((g.r.Int()%(cb+cb))-cb) / cb
		}
		normalize3(&g.g3[i])
	}

	for i := cb - 1; i > 0; i-- {
		k := g.p[i]
		j := g.r.Int() % cb
		g.p[i] = g.p[j]
		g.p[j] = k
	}

	for i := 0; i < cb+2; i++ {
		g.p[cb+i] = g.p[i]
		g.g1[cb+i] = g.g1[i]
		for j := 0; j < 2; j++ {
			g.g2[cb+i][j] = g.g2[i][j]
		}
		for j := 0; j < 3; j++ {
			g.g3[cb+i][j] = g.g3[i][j]
		}
	}
}

// Noise1D calculates Perlin noise at point x.
func (g *Generator) Noise1D(x float64) float64 {
	var scale float64 = 1
	var sum float64 = 0
	p := x

	for i := 0; i < g.N; i++ {
		val := g.noise1(p)
		sum += val / scale
		scale *= g.Alpha
		p *= g.Beta
	}
	return sum
}

// Noise2D calculates Perlin noise at point x,y.
func (g *Generator) Noise2D(x, y float64) float64 {
	var scale float64 = 1
	var sum float64 = 0
	var p [2]float64

	p[0] = x
	p[1] = y

	for i := 0; i < g.N; i++ {
		val := g.noise2(p)
		sum += val / scale
		scale *= g.Alpha
		p[0] *= g.Beta
		p[1] *= g.Beta
	}
	return sum
}

// Noise3D calculates Perlin noise at point x,y,z.
// For small values of z this will be equivalent to calling Noise2D(x, y).
func (g *Generator) Noise3D(x, y, z float64) float64 {
	if z < 0.0001 {
		return g.Noise2D(x, y)
	}

	var scale float64 = 1
	var sum float64 = 0
	var p [3]float64

	p[0] = x
	p[1] = y
	p[2] = z

	for i := 0; i < g.N; i++ {
		val := g.noise3(p)
		sum += val / scale
		scale *= g.Alpha
		p[0] *= g.Beta
		p[1] *= g.Beta
		p[2] *= g.Beta
	}
	return sum
}

func (g *Generator) noise1(arg float64) float64 {
	var vec [1]float64
	vec[0] = arg

	t := vec[0] + cn
	bx0 := int(t) & cbm
	bx1 := (bx0 + 1) & cbm
	rx0 := t - float64(int(t))
	rx1 := rx0 - 1.

	sx := sCurve(rx0)
	u := rx0 * g.g1[g.p[bx0]]
	v := rx1 * g.g1[g.p[bx1]]

	return lerp(sx, u, v)
}

func (g *Generator) noise2(vec [2]float64) float64 {
	t := vec[0] + cn
	bx0 := int(t) & cbm
	bx1 := (bx0 + 1) & cbm
	rx0 := t - float64(int(t))
	rx1 := rx0 - 1.

	t = vec[1] + cn
	by0 := int(t) & cbm
	by1 := (by0 + 1) & cbm
	ry0 := t - float64(int(t))
	ry1 := ry0 - 1.

	i := g.p[bx0]
	j := g.p[bx1]

	b00 := g.p[i+by0]
	b10 := g.p[j+by0]
	b01 := g.p[i+by1]
	b11 := g.p[j+by1]

	sx := sCurve(rx0)
	sy := sCurve(ry0)

	q := g.g2[b00]
	u := at2(rx0, ry0, q)
	q = g.g2[b10]
	v := at2(rx1, ry0, q)
	a := lerp(sx, u, v)

	q = g.g2[b01]
	u = at2(rx0, ry1, q)
	q = g.g2[b11]
	v = at2(rx1, ry1, q)
	b := lerp(sx, u, v)

	return lerp(sy, a, b)
}

func (g *Generator) noise3(vec [3]float64) float64 {
	t := vec[0] + cn
	bx0 := int(t) & cbm
	bx1 := (bx0 + 1) & cbm
	rx0 := t - float64(int(t))
	rx1 := rx0 - 1.

	t = vec[1] + cn
	by0 := int(t) & cbm
	by1 := (by0 + 1) & cbm
	ry0 := t - float64(int(t))
	ry1 := ry0 - 1.

	t = vec[2] + cn
	bz0 := int(t) & cbm
	bz1 := (bz0 + 1) & cbm
	rz0 := t - float64(int(t))
	rz1 := rz0 - 1.

	i := g.p[bx0]
	j := g.p[bx1]

	b00 := g.p[i+by0]
	b10 := g.p[j+by0]
	b01 := g.p[i+by1]
	b11 := g.p[j+by1]

	t = sCurve(rx0)
	sy := sCurve(ry0)
	sz := sCurve(rz0)

	q := g.g3[b00+bz0]
	u := at3(rx0, ry0, rz0, q)
	q = g.g3[b10+bz0]
	v := at3(rx1, ry0, rz0, q)
	a := lerp(t, u, v)

	q = g.g3[b01+bz0]
	u = at3(rx0, ry1, rz0, q)
	q = g.g3[b11+bz0]
	v = at3(rx1, ry1, rz0, q)
	b := lerp(t, u, v)

	c := lerp(sy, a, b)

	q = g.g3[b00+bz1]
	u = at3(rx0, ry0, rz1, q)
	q = g.g3[b10+bz1]
	v = at3(rx1, ry0, rz1, q)
	a = lerp(t, u, v)

	q = g.g3[b01+bz1]
	u = at3(rx0, ry1, rz1, q)
	q = g.g3[b11+bz1]
	v = at3(rx1, ry1, rz1, q)
	b = lerp(t, u, v)

	d := lerp(sy, a, b)

	return lerp(sz, c, d)
}

func normalize2(v *[2]float64) {
	s := math.Sqrt(v[0]*v[0] + v[1]*v[1])
	v[0] = v[0] / s
	v[1] = v[1] / s
}

func normalize3(v *[3]float64) {
	s := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	v[0] = v[0] / s
	v[1] = v[1] / s
	v[2] = v[2] / s
}

func at2(rx, ry float64, q [2]float64) float64 {
	return rx*q[0] + ry*q[1]
}

func at3(rx, ry, rz float64, q [3]float64) float64 {
	return rx*q[0] + ry*q[1] + rz*q[2]
}

func sCurve(t float64) float64 {
	return t * t * (3. - 2.*t)
}

func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}
