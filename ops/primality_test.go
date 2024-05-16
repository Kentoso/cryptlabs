package ops

import (
	"math/big"
	"testing"
)

func TestMillerRabin(t *testing.T) {
	tests := []struct {
		name    string
		n       *big.Int
		k       int
		isPrime bool
	}{
		{
			name:    "Test with small prime number",
			n:       big.NewInt(13),
			k:       5,
			isPrime: true,
		},
		{
			name:    "Test with small non-prime number",
			n:       big.NewInt(12),
			k:       5,
			isPrime: false,
		},
		{
			name:    "Test with large prime number",
			n:       big.NewInt(982451653), // Large known prime
			k:       10,
			isPrime: true,
		},
		{
			name:    "Test with large non-prime number",
			n:       big.NewInt(982451654),
			k:       10,
			isPrime: false,
		},
		{
			name:    "Test with very large non-prime number",
			n:       new(big.Int).Mod(big.NewInt(982451654), big.NewInt(982451654)),
			k:       20,
			isPrime: false,
		},
		{
			name:    "Test with two (edge case)",
			n:       big.NewInt(2),
			k:       5,
			isPrime: true,
		},
		{
			name:    "Test with one (edge case)",
			n:       big.NewInt(1),
			k:       5,
			isPrime: false,
		},
		{
			name:    "Test with zero (edge case)",
			n:       big.NewInt(0),
			k:       5,
			isPrime: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MillerRabin(tc.n, tc.k); got != tc.isPrime {
				t.Errorf("MillerRabin(%s, %d): expected %t, got %t", tc.n, tc.k, tc.isPrime, got)
			}
		})
	}
}
