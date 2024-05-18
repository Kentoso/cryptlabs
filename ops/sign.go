package ops

import "math/big"

func Sign(priv PrivateKey, hash byte) *big.Int {
	h := big.NewInt(int64(hash))
	signature := PowMod(h, priv.D, priv.N)
	return signature
}

func Verify(pub PublicKey, hash byte, signature *big.Int) bool {
	h := big.NewInt(int64(hash))
	verifiedHash := PowMod(signature, pub.E, pub.N)
	return h.Cmp(verifiedHash) == 0
}
