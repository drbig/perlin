// See LICENSE.txt for licensing information.

package perlin

import (
	"fmt"
	"testing"
)

var (
	g = NewGenerator(2, 2, 1, 100)
)

func TestNoise1D(t *testing.T) {
	noise := g.Noise1D(0.2)
	if noise < -1 || noise > 1 {
		t.Errorf("Unexpected noise")
	}
}

func TestNoise2D(t *testing.T) {
	noise := g.Noise2D(0.2, 0.2)
	if noise < -1 || noise > 1 {
		t.Errorf("Unexpected noise")
	}
}

func TestNoise3D(t *testing.T) {
	noise := g.Noise3D(0.2, 0.2, 0.2)
	if noise < -1 || noise > 1 {
		t.Errorf("Unexpected noise")
	}
}

func TestTwoGenerators(t *testing.T) {
	g2 := NewGenerator(2, 2, 1, 41)
	val := g.Noise1D(0.2)
	val2 := g2.Noise1D(0.2)
	if val == val2 {
		t.Errorf("Same value for Noise1D")
	}
	val = g.Noise2D(0.12, 0.48)
	val2 = g2.Noise2D(0.12, 0.48)
	if val == val2 {
		t.Errorf("Same value for Noise2D")
	}
	val = g.Noise3D(0.19, 9.32, 0.1233)
	val2 = g2.Noise3D(0.19, 0.32, 0.1233)
	if val == val2 {
		t.Errorf("Same value for Noise3D")
	}
}

func TestDeterministicOutput(t *testing.T) {
	val := g.Noise2D(0.412, 0.931)
	val2 := g.Noise2D(0.412, 0.931)
	if val != val2 {
		t.Errorf("Non-deterministic output: %v != %v", val, val2)
	}
}

func TestLowZNoise3D(t *testing.T) {
	val := g.Noise3D(0.23, 0.23, 0.0000012)
	val2 := g.Noise2D(0.23, 0.23)
	if val != val2 {
		t.Errorf("Noise3D mismatches Noise2D at near-zero Z: %v != %v", val, val2)
	}
}

func ExampleGenerator_Noise1D() {
	g1 := NewGenerator(2, 2, 1, 41)
	g2 := NewGenerator(2, 2, 1, 4231)
	fmt.Printf("g1 noise at 0.1, 0.2 = %f\n", g1.Noise2D(0.1, 0.2))
	fmt.Printf("g2 noise at 0.1, 0.2 = %f\n", g2.Noise2D(0.1, 0.2))
	// Output:
	// g1 noise at 0.1, 0.2 = -0.125053
	// g2 noise at 0.1, 0.2 = 0.182675
}

func BenchmarkNoise1D(b *testing.B) {
	step := 1.0 / float64(b.N)
	for i := float64(0); i < 1.0; i += step {
		g.Noise1D(i)
	}
}

func BenchmarkNoise2D(b *testing.B) {
	step := 1.0 / float64(b.N)
	for i := float64(0); i < 1.0; i += step {
		g.Noise2D(i, i)
	}
}

func BenchmarkNoise3D(b *testing.B) {
	step := 1.0 / float64(b.N)
	for i := float64(0); i < 1.0; i += step {
		g.Noise3D(i, i, i)
	}
}
