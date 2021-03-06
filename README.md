# perlin [![Build Status](https://travis-ci.org/drbig/perlin.svg?branch=master)](https://travis-ci.org/drbig/perlin) [![Coverage Status](https://img.shields.io/coveralls/drbig/perlin.svg)](https://coveralls.io/r/drbig/perlin?branch=master) [![GoDoc](https://godoc.org/github.com/drbig/perlin?status.svg)](http://godoc.org/github.com/drbig/perlin)

Yet another Go library implementing Perlin noise.

Features:

- You can have multiple independent generators
- Each generator can be re-parametrised, re-seeded and reused
- Noise for 1, 2 and 3 dimensions
- Full test coverage, includes benchmarks

The library is a direct rip-off of [aquilax/go-perlin](https://github.com/aquilax/go-perlin), which in turn has been based on [this](https://git.gnome.org/browse/gegl/tree/operations/common/perlin), which was ultimately based on the work of [Ken Perlin](http://en.wikipedia.org/wiki/Ken_Perlin). I'm licensing it under my usual two clause BSD license, if anybody has a problem with that please do tell.

## Showcase

Using the included `perlin` command:

    $ cd cmd/perlin
    $ go build
    $ ./perlin
    Usage: ./perlin [options] path
      -a=2: alpha factor
      -b=2: beta factor
      -c=false: smooth (continuous) gradient
      -h=256: image height
      -n=1: octave factor
      -r=1417875091: random seed
      -s=1: scaling factor
      -w=320: image width
    $ ./perlin -c -a 1.25 -b 1.88 -n 10 -r 1 example.png

![Example output](https://raw.github.com/drbig/perlin/master/example.png)

## Usage notes

All functions return values in the range of `<-1.0, 1.0>`. Once seeded the generator is fully deterministic. Please see docs for more information.

Benchmarks (linux x64, Intel i7-2620M @ 2.70GHz):

    PASS
    BenchmarkNoise1D        100000000               14.9 ns/op
    BenchmarkNoise2D        100000000               27.0 ns/op
    BenchmarkNoise3D        50000000                46.6 ns/op
    ok      github.com/drbig/perlin 6.631s

## Licensing

Standard two-clause BSD license, see LICENSE.txt for details.

Copyright (c) 2014 Piotr S. Staszewski
