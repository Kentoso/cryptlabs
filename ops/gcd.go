package ops

import (
	"github.com/Kentoso/cryptlabs/constants"
	"math/big"
)

func ExtendedGCD(a, b *big.Int) (gcd *big.Int, x *big.Int, y *big.Int) {
	x = big.NewInt(0)
	y = big.NewInt(1)
	lastX := big.NewInt(1)
	lastY := big.NewInt(0)

	aa := new(big.Int).Set(a)
	bb := new(big.Int).Set(b)

	for bb.Cmp(constants.BigZero) != 0 {
		q := new(big.Int).Div(aa, bb)
		r := new(big.Int).Mod(aa, bb)

		aa.Set(bb)
		bb.Set(r)

		tempX := new(big.Int).Set(x)
		x.Mul(q, x).Sub(lastX, x)
		lastX.Set(tempX)

		tempY := new(big.Int).Set(y)
		y.Mul(q, y).Sub(lastY, y)
		lastY.Set(tempY)
	}

	return aa, lastX, lastY
}

func ModInverse(a, m *big.Int) *big.Int {
	gcd, x, _ := ExtendedGCD(a, m)
	if gcd.Cmp(big.NewInt(1)) != 0 {
		return nil
	}
	if x.Cmp(big.NewInt(0)) < 0 {
		x.Add(x, m)
	}
	return x
}
