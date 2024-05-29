package main

import (
	"fmt"
	"testing"
)

// BenchmarkScenario
// BenchmarkScenario/128_bits
// BenchmarkScenario/128_bits-8         	      72	  15660702 ns/op
// BenchmarkScenario/256_bits
// BenchmarkScenario/256_bits-8         	      24	  54681012 ns/op
// BenchmarkScenario/384_bits
// BenchmarkScenario/384_bits-8         	       8	 184788745 ns/op
// BenchmarkScenario/512_bits
// BenchmarkScenario/512_bits-8         	       4	 344028812 ns/op
// BenchmarkScenario/640_bits
// BenchmarkScenario/640_bits-8         	       3	 465462375 ns/op
// BenchmarkScenario/768_bits
// BenchmarkScenario/768_bits-8         	       2	 855871854 ns/op
// BenchmarkScenario/896_bits
// BenchmarkScenario/896_bits-8         	       1	1051060542 ns/op
// BenchmarkScenario/1024_bits
// BenchmarkScenario/1024_bits-8        	       1	2764639334 ns/op
// BenchmarkScenario/1152_bits
// BenchmarkScenario/1152_bits-8        	       1	2098315500 ns/op
// BenchmarkScenario/1280_bits
// BenchmarkScenario/1280_bits-8        	       1	5787372708 ns/op
// BenchmarkScenario/2048_bits
// BenchmarkScenario/2048_bits-8        	       1	28629616125 ns/op
// PASS

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
