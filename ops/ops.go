package ops

import (
	"github.com/Kentoso/cryptlabs/constants"
	"math/big"
)

func MulMod(x, y, m *big.Int) *big.Int {
	xMod, yMod := new(big.Int).Set(x), new(big.Int).Set(y)
	if xMod.Cmp(m) > 0 {
		xMod.Mod(x, m)
	}
	if yMod.Cmp(m) > 0 {
		yMod.Mod(y, m)
	}

	return xMod.Mul(xMod, yMod).Mod(xMod, m)
}

func PowMod(base, exponent, m *big.Int) *big.Int {
	if m.Cmp(constants.BigOne) == 0 {
		return constants.BigZero
	}

	result := big.NewInt(1)

	tempBase := new(big.Int).Mod(base, m)

	for i := 0; i < exponent.BitLen(); i++ {
		if exponent.Bit(i) == 1 {
			result = MulMod(result, tempBase, m)
		}
		tempBase = MulMod(tempBase, tempBase, m)
	}

	return result
}
