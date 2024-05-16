package ops

import (
	"github.com/Kentoso/cryptlabs/constants"
	"math/big"
	"math/rand"
	"time"
)

type PrimalityTest func(n *big.Int, k int) bool

func MillerRabin(n *big.Int, k int) bool {
	if n.Cmp(constants.BigTwo) == -1 {
		return false
	}
	if n.Cmp(constants.BigTwo) == 0 {
		return true
	}
	if new(big.Int).And(n, constants.BigOne).Cmp(constants.BigZero) == 0 {
		return false
	}

	s := big.NewInt(0)
	d := new(big.Int).Sub(n, constants.BigOne)

	// looking at least significant bit of d, if 0 - even else odd
	// d % 2 == 0
	for new(big.Int).And(d, constants.BigOne).Cmp(constants.BigZero) == 0 {
		d.Rsh(d, 1)
		s.Add(s, constants.BigOne)
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < k; i++ {
		a := randBigInt(randSource, new(big.Int).Sub(n, constants.BigTwo))
		a.Add(a, constants.BigTwo)

		// a^d mod n
		x := PowMod(a, d, n)
		if x.Cmp(constants.BigOne) == 0 || x.Cmp(new(big.Int).Sub(n, constants.BigOne)) == 0 {
			continue
		}
		for j := big.NewInt(0); j.Cmp(new(big.Int).Sub(s, constants.BigOne)) == -1; j.Add(j, constants.BigOne) {
			x = MulMod(x, x, n)
			if x.Cmp(new(big.Int).Sub(n, constants.BigOne)) == 0 {
				break
			}
		}
		if x.Cmp(new(big.Int).Sub(n, constants.BigOne)) != 0 {
			return false
		}
	}
	return true
}

func randBigInt(rnd *rand.Rand, n *big.Int) *big.Int {
	result := new(big.Int)
	return result.Rand(rnd, n)
}
