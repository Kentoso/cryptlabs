package ops

import (
	"fmt"
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

func BailliePSW(n *big.Int, _ int) bool {
	if n.Cmp(constants.BigTwo) == -1 {
		return false
	}
	if n.Cmp(constants.BigTwo) == 0 {
		return true
	}
	if new(big.Int).And(n, constants.BigOne).Cmp(constants.BigZero) == 0 {
		return false
	}

	if !MillerRabinBase2(n) {
		return false
	}

	if !LucasTest(n) {
		return false
	}

	return true
}

func MillerRabinBase2(n *big.Int) bool {
	if n.Cmp(constants.BigTwo) < 0 {
		return false
	}
	if n.Cmp(constants.BigTwo) == 0 {
		return true
	}
	if new(big.Int).And(n, constants.BigOne).Cmp(constants.BigZero) == 0 {
		return false
	}

	d := new(big.Int).Sub(n, constants.BigOne)
	s := big.NewInt(0)

	for new(big.Int).And(d, constants.BigOne).Cmp(constants.BigZero) == 0 {
		d.Rsh(d, 1)
		s.Add(s, constants.BigOne)
	}

	a := constants.BigTwo
	x := new(big.Int).Exp(a, d, n)

	if x.Cmp(constants.BigOne) == 0 || x.Cmp(new(big.Int).Sub(n, constants.BigOne)) == 0 {
		return true
	}

	for s.Cmp(constants.BigOne) > 0 {
		x = MulMod(x, x, n)
		if x.Cmp(new(big.Int).Sub(n, constants.BigOne)) == 0 {
			return true
		}
		s.Sub(s, constants.BigOne)
	}

	return false
}

func LucasTest(n *big.Int) bool {
	d := big.NewInt(5)
	isNegative := false
	for {
		jacobi := Jacobi(d, n)
		if jacobi == -1 {
			break
		}
		if isNegative {
			d.Sub(d, big.NewInt(2))
		} else {
			d.Add(d, big.NewInt(2))
		}
		d.Neg(d)
		isNegative = !isNegative
		fmt.Println(d)
	}
	return LucasProbablePrime(n, d)
}

func Jacobi(a, n *big.Int) int {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	eight := big.NewInt(8)
	four := big.NewInt(4)
	three := big.NewInt(3)
	five := big.NewInt(5)
	seven := big.NewInt(7)

	// Step 1: a = a mod n
	a.Mod(a, n)

	// Step 2: If a = 1 or n = 1, then return 1
	if a.Cmp(one) == 0 || n.Cmp(one) == 0 {
		return 1
	}

	// Step 3: If a = 0, then return 0
	if a.Cmp(zero) == 0 {
		return 0
	}

	// Step 4: Define e and a1 such that a = 2^e * a1, where a1 is odd
	e := 0
	a1 := new(big.Int).Set(a)
	for a1.Bit(0) == 0 { // While a1 is even
		a1.Rsh(a1, 1)
		e++
	}

	// Step 5: Determine s based on e
	s := 1
	if e%2 == 1 {
		nMod8 := new(big.Int).Mod(n, eight)
		if nMod8.Cmp(one) == 0 || nMod8.Cmp(seven) == 0 {
			s = 1
		} else if nMod8.Cmp(three) == 0 || nMod8.Cmp(five) == 0 {
			s = -1
		}
	}

	// Step 6: Adjust s based on n mod 4 and a1 mod 4
	nMod4 := new(big.Int).Mod(n, four)
	a1Mod4 := new(big.Int).Mod(a1, four)
	if nMod4.Cmp(three) == 0 && a1Mod4.Cmp(three) == 0 {
		s = -s
	}

	// Step 7: n1 = n mod a1
	n1 := new(big.Int).Mod(n, a1)

	// Step 8: Return s * Jacobi(n1, a1) (recursive call)
	return s * Jacobi(n1, a1)
}

func LucasProbablePrime(n, d *big.Int) bool {
	one := big.NewInt(1)
	two := big.NewInt(2)
	zero := big.NewInt(0)

	// Step 1: Check if n is a perfect square
	root := new(big.Int).Sqrt(n)
	if root.Mul(root, root).Cmp(n) == 0 {
		return false // n is a perfect square, return COMPOSITE
	}

	// Step 3: K = n + 1
	k := new(big.Int).Add(n, one)

	// Step 4: Binary expansion of K
	kBits := k.BitLen()

	// Step 5: Initialize U_r = 1 and V_r = 1
	u := big.NewInt(1)
	v := big.NewInt(1)

	// Helper function to handle A/2 mod C
	halfMod := func(a, c *big.Int) *big.Int {
		if new(big.Int).Mod(a, two).Cmp(zero) != 0 { // if a is odd
			half := new(big.Int).Add(a, c)
			return half.Rsh(half, 1).Mod(half, c)
		}
		half := new(big.Int).Rsh(a, 1)
		return half.Mod(half, c)
	}

	// Step 6: Loop over the binary representation of K
	for i := kBits - 2; i >= 0; i-- {
		// Step 6.1: U_temp = U_i * V_i mod n
		uTemp := new(big.Int).Mul(u, v)
		uTemp.Mod(uTemp, n)

		// Step 6.2: V_temp = (V_i^2 + d * U_i^2) / 2 mod n
		vTemp := new(big.Int).Mul(v, v)
		vTemp.Add(vTemp, new(big.Int).Mul(d, new(big.Int).Mul(u, u)))
		vTemp = halfMod(vTemp, n)

		// Step 6.3: Conditional update based on K_i
		if k.Bit(i) == 1 {
			// If K_i = 1
			u = halfMod(new(big.Int).Add(uTemp, vTemp), n)

			v = halfMod(new(big.Int).Add(vTemp, new(big.Int).Mul(d, uTemp)), n)
		} else {
			// If K_i = 0
			u = uTemp
			v = vTemp
		}
	}

	// Step 7: If U_0 = 0, return PROBABLY PRIME. Otherwise, return COMPOSITE.
	return u.Cmp(zero) == 0
}
