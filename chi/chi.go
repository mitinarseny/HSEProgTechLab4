package chi

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func IsUniform(sample []uint32, signLevel float64) bool {
	if signLevel <= 0 || signLevel >= 1 {
		panic("significance level should be between 0 and 1")
	}
	chi := chiSquared(sample)
	return chi <= distuv.ChiSquared{
		K:   float64(len(sample) - 1),
	}.Quantile(1-signLevel)
}

func chiSquared(sample []uint32) float64 {
	k := 1 + int(math.Log2(float64(len(sample))))

	intervalEdges := make([]uint32, k+1)
	intervalEdges[0] = 0
	for i := 1; i < k; i++ {
		intervalEdges[i] = ^uint32(0) / uint32(k) * uint32(i)
	}
	intervalEdges[k] = ^uint32(0)

	intervalCounts := make([]int, 0, k)

	for i := 0; i < k; i++ {
		var count int
		for _, s := range sample {
			if intervalEdges[i] < s && s <= intervalEdges[i+1] {
				count++
			}
		}
		intervalCounts = append(intervalCounts, count)
	}

	theorPi := 1 / float64(k)

	var chi float64
	for _, n := range intervalCounts {
		chi += float64(n*n) / theorPi
	}
	chi /= float64(len(sample))
	chi -= float64(len(sample))
	return chi
}
