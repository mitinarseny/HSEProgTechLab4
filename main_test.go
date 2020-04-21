package HSEProgTechLab4

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/mitinarseny/HSEProgTechLab4/chi"
	"github.com/mitinarseny/HSEProgTechLab4/custom"
	"github.com/mitinarseny/HSEProgTechLab4/dummy"
)

type Generator interface {
	Gen() uint32
}

const (
	N    = 10
	SIZE = 1000
)

func TestDummy(t *testing.T) {
	test(t, N, SIZE, func() Generator {
		return dummy.New(rand.Uint32())
	})
}

func TestCustom(t *testing.T) {
	test(t, N, SIZE, func() Generator {
		return custom.New(^uint32(0), 1220703125, 7, rand.Uint32())
	})
}

func test(t *testing.T, N, eachSize int, makeGenerator func() Generator) {
	samples := genNSamples(N, eachSize, makeGenerator)
	for i, sample := range samples {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			m := mean(sample)
			t.Logf("Mean:\t%f", m)
			D := deviation(sample, m)
			sigma := math.Sqrt(D)
			t.Logf("Sqrt deviation:\t%f", sigma)
			c := kVar(sigma, m)
			t.Logf("Variance coefficient:\t%f", c)
			t.Logf("Is Uniform: %t", chi.IsUniform(sample, 0.05))
		})
	}
}

func mean(sample []uint32) float64 {
	var sum uint32
	for _, i := range sample {
		sum += i
	}
	return float64(sum) / float64(len(sample))
}

func deviation(sample []uint32, mean float64) float64 {
	var D float64
	for _, i := range sample {
		delta := float64(i) - mean
		D += delta * delta
	}
	D /= float64(len(sample))
	return D
}

func kVar(sigma, mean float64) float64 {
	return sigma / mean
}

func genNSamples(N, eachSize int, makeGenerator func() Generator) [][]uint32 {
	s := make([][]uint32, 0, N)
	for i := 0; i < N; i++ {
		s = append(s, genSample(eachSize, makeGenerator()))
	}
	return s
}

func genSample(size int, g Generator) []uint32 {
	s := make([]uint32, 0, size)
	for i := 0; i < size; i++ {
		s = append(s, g.Gen())
	}
	return s
}

func BenchmarkDummy(b *testing.B) {
	benchmark(b, dummy.New(0).Gen)
}

func BenchmarkCustom(b *testing.B) {
	benchmark(b, custom.New(^uint32(0), 1220703125, 7, 7).Gen)
}
func BenchmarkStd(b *testing.B) {
	benchmark(b, rand.New(rand.NewSource(time.Now().UnixNano())).Uint32)
}

func benchmark(b *testing.B, generator func() uint32) {
	sizes := []int{
		1e3,
		5e3,
		1e4,
		5e4,
		1e5,
		5e5,
		1e6,
	}
	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < size; j++ {
					generator()
				}
			}
		})
	}
}
