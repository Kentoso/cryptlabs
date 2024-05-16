package ops

import "math/big"

func LCM(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)
	if a.Cmp(zero) == 0 || b.Cmp(zero) == 0 {
		return big.NewInt(0)
	}

	gcd, _, _ := ExtendedGCD(a, b)

	product := new(big.Int).Mul(a, b)
	product = product.Abs(product)

	lcm := new(big.Int).Div(product, gcd)

	return lcm
}
