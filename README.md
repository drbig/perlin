# perlin [![Build Status](https://travis-ci.org/drbig/perlin.svg?branch=master)](https://travis-ci.org/drbig/perlin) [![Coverage Status](https://img.shields.io/coveralls/drbig/perlin.svg)](https://coveralls.io/r/drbig/perlin?branch=master) [![GoDoc](https://godoc.org/github.com/drbig/perlin?status.svg)](http://godoc.org/github.com/drbig/perlin)

Yet another Go library implementing Perlin noise.

Features:

- You can have multiple independent generators
- Each generator can be re-parametrised, re-seeded and reused
- Noise for 1, 2 and 3 dimensions
- Full test coverage, includes benchmarks

This is a direct rip-off of [aquilax/go-perlin](https://github.com/aquilax/go-perlin), which in turn has been based on [this](https://git.gnome.org/browse/gegl/tree/operations/common/perlin), which was ultimately based on the work of [Ken Perlin](http://en.wikipedia.org/wiki/Ken_Perlin). I'm licensing it under my usual two clause BSD license, if anybody has a problem with that please do tell.

## Usage notes

All functions return values in the range of `<-1, 1>`, and it seems that they require the input coordinates in the same range. Once seeded the generator is fully deterministic. Please see docs for more information.

Benchmarks (linux x64, Intel i7-2620M @ 2.70GHz):

    PASS
    BenchmarkNoise1D        100000000               14.9 ns/op
    BenchmarkNoise2D        100000000               27.0 ns/op
    BenchmarkNoise3D        50000000                46.6 ns/op
    ok      github.com/drbig/perlin 6.631s

## Licensing

Standard two-clause BSD license, see LICENSE.txt for details.

Copyright (c) 2014 Piotr S. Staszewski
