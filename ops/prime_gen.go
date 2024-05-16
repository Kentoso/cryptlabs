package ops

import (
	"math/big"
	"math/rand"
	"time"
)

func FindNBitPrime(primalityTest func(n *big.Int, k int) bool, n int, k int) *big.Int {
	if n <= 1 {
		return nil
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	min := new(big.Int).Lsh(big.NewInt(1), uint(n-1))
	max := new(big.Int).Lsh(big.NewInt(1), uint(n))
	prime := new(big.Int)

	for {
		prime.Rand(rng, new(big.Int).Sub(max, min)).Add(prime, min)
		prime.SetBit(prime, int(n-1), 1) // n bits
		prime.SetBit(prime, 0, 1)        // odd

		if primalityTest(prime, k) {
			return prime
		}
	}
}
