package mandelbrot

import (
	"runtime"
	"testing"
)

func BenchmarkSerialRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SerialRender()
	}
}

func benchConcurrentRender(b *testing.B, procs int) {
	for i := 0; i < b.N; i++ {
		ConcurrentRender(procs)
	}
}

func Benchmark1(b *testing.B) {
	benchConcurrentRender(b, 1)
}

func BenchmarkMaxProces(b *testing.B) {
	benchConcurrentRender(b, runtime.GOMAXPROCS(-1))
}

func Benchmark8(b *testing.B) {
	benchConcurrentRender(b, 8)
}

func Benchmark16(b *testing.B) {
	benchConcurrentRender(b, 16)
}

func Benchmark32(b *testing.B) {
	benchConcurrentRender(b, 32)
}

func Benchmark64(b *testing.B) {
	benchConcurrentRender(b, 64)
}

func Benchmark128(b *testing.B) {
	benchConcurrentRender(b, 128)
}