package ops

import (
	"github.com/Kentoso/cryptlabs/constants"
	"math/big"
)

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	N *big.Int
	D *big.Int
}

func RSA(bit int, primalityTest PrimalityTest) (PublicKey, PrivateKey) {
	var p, q *big.Int
	for p.Cmp(q) == 0 {
		p, q = FindNBitPrime(primalityTest, bit, 20), FindNBitPrime(primalityTest, bit, 20)
	}
	n := new(big.Int).Mul(p, q)

	carmichael := LCM(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	commonEs := []*big.Int{
		big.NewInt(65537),
		big.NewInt(257),
		big.NewInt(17),
		big.NewInt(5),
		big.NewInt(3),
	}

	var e *big.Int
	for _, currE := range commonEs {
		if currE.Cmp(carmichael) == -1 {
			gcd, _, _ := ExtendedGCD(currE, carmichael)
			if gcd.Cmp(constants.BigOne) == 0 {
				e = currE
				break
			}
		}
	}

	_, d, _ := ExtendedGCD(e, carmichael)
	if d.Cmp(constants.BigZero) == -1 {
		d.Add(d, carmichael)
	}

	//fmt.Println(p, q, n, carmichael, e, d)

	return PublicKey{N: n, E: e}, PrivateKey{N: n, D: d}
}

func Encrypt(publicKey PublicKey, message string) *big.Int {
	return PowMod(textToBigInt(message), publicKey.E, publicKey.N)
}

func Decrypt(privateKey PrivateKey, cipher *big.Int) string {
	return bigIntToText(PowMod(cipher, privateKey.D, privateKey.N))
}

func textToBigInt(message string) *big.Int {
	return new(big.Int).SetBytes([]byte(message))
}

func bigIntToText(bi *big.Int) string {
	return string(bi.Bytes())
}
