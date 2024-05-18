package main

import (
	"fmt"
	"testing"
)

func BenchmarkScenario(b *testing.B) {
	bitCounts := []int{128, 256, 384, 512, 640, 768, 896, 1024, 1152, 1280, 2048}
	for _, bits := range bitCounts {
		b.Run(fmt.Sprintf("%d bits", bits), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				scenario(bits, false)
			}
		})
	}
}
