package ops

import (
	"math/big"
	"testing"
)

func TestExtendedGCD(t *testing.T) {
	tests := []struct {
		a, b        *big.Int
		expectedGCD *big.Int
		expectedX   *big.Int
		expectedY   *big.Int
	}{
		{
			a: big.NewInt(240), b: big.NewInt(46),
			expectedGCD: big.NewInt(2), expectedX: big.NewInt(-9), expectedY: big.NewInt(47),
		},
		{
			a: big.NewInt(0), b: big.NewInt(0),
			expectedGCD: big.NewInt(0), expectedX: big.NewInt(0), expectedY: big.NewInt(1),
		},
		{
			a: big.NewInt(1), b: big.NewInt(0),
			expectedGCD: big.NewInt(1), expectedX: big.NewInt(1), expectedY: big.NewInt(0),
		},
		{
			a: big.NewInt(0), b: big.NewInt(1),
			expectedGCD: big.NewInt(1), expectedX: big.NewInt(0), expectedY: big.NewInt(1),
		},
		{
			a: big.NewInt(17), b: big.NewInt(12),
			expectedGCD: big.NewInt(1), expectedX: big.NewInt(1), expectedY: big.NewInt(-1),
		},
		{
			a: big.NewInt(120), b: big.NewInt(23),
			expectedGCD: big.NewInt(1), expectedX: big.NewInt(-9), expectedY: big.NewInt(47),
		},
	}

	for _, test := range tests {
		gcd, x, y := ExtendedGCD(test.a, test.b)

		if gcd.Cmp(test.expectedGCD) != 0 {
			t.Errorf("GCD(%s, %s) = %s; want %s", test.a, test.b, gcd, test.expectedGCD)
		}

		// Verify the linear combination a*x + b*y == gcd
		ax := new(big.Int).Mul(test.a, x)
		by := new(big.Int).Mul(test.b, y)
		sum := new(big.Int).Add(ax, by)

		if sum.Cmp(gcd) != 0 {
			t.Errorf("%s * %s + %s * %s = %s; want %s", test.a, x, test.b, y, sum, gcd)
		}
	}
}
