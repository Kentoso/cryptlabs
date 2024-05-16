package ops

import (
	"math/big"
	"testing"
)

func TestMulMod(t *testing.T) {
	tests := []struct {
		name   string
		x      *big.Int
		y      *big.Int
		m      *big.Int
		expect *big.Int
	}{
		{
			name:   "Simple multiplication",
			x:      big.NewInt(6),
			y:      big.NewInt(7),
			m:      big.NewInt(10),
			expect: big.NewInt(2),
		},
		{
			name:   "Multiplication with zero",
			x:      big.NewInt(0),
			y:      big.NewInt(7),
			m:      big.NewInt(10),
			expect: big.NewInt(0),
		},
		{
			name:   "Large numbers",
			x:      big.NewInt(123456),
			y:      big.NewInt(654321),
			m:      big.NewInt(1000),
			expect: big.NewInt(376),
		},
		{
			name:   "Modulo one",
			x:      big.NewInt(500),
			y:      big.NewInt(300),
			m:      big.NewInt(1),
			expect: big.NewInt(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := MulMod(tc.x, tc.y, tc.m)
			if result.Cmp(tc.expect) != 0 {
				t.Errorf("MulMod(%d, %d, %d): expected %d, got %d", tc.x, tc.y, tc.m, tc.expect, result)
			}
		})
	}
}

func TestPowMod(t *testing.T) {
	tests := []struct {
		name     string
		base     *big.Int
		exponent *big.Int
		m        *big.Int
		expect   *big.Int
	}{
		{
			name:     "Simple exponentiation",
			base:     big.NewInt(2),
			exponent: big.NewInt(3),
			m:        big.NewInt(5),
			expect:   big.NewInt(3),
		},
		{
			name:     "Zero exponent",
			base:     big.NewInt(4),
			exponent: big.NewInt(0),
			m:        big.NewInt(10),
			expect:   big.NewInt(1),
		},
		{
			name:     "Modulo one",
			base:     big.NewInt(2),
			exponent: big.NewInt(3),
			m:        big.NewInt(1),
			expect:   big.NewInt(0),
		},
		{
			name:     "Large exponent",
			base:     big.NewInt(2),
			exponent: big.NewInt(20),
			m:        big.NewInt(17),
			expect:   big.NewInt(16),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := PowMod(tc.base, tc.exponent, tc.m)
			if result.Cmp(tc.expect) != 0 {
				t.Errorf("PowMod(%d, %d, %d): expected %d, got %d", tc.base, tc.exponent, tc.m, tc.expect, result)
			}
		})
	}
}
