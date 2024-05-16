package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Kentoso/cryptlabs/ops"
)

func main() {
	probablyNBitPrime := ops.FindNBitPrime(ops.MillerRabin, 500, 20)

	fmt.Printf("Base 10: %d\n", probablyNBitPrime) // Default String() method outputs base 10

	binary := probablyNBitPrime.Text(2)
	fmt.Printf("Base 2: %s, %d bits\n", binary, len(binary))

	bytes := probablyNBitPrime.Bytes()

	base64String := base64.StdEncoding.EncodeToString(bytes)
	fmt.Printf("Base 64: %s\n", base64String)

	fmt.Printf("Byte Array: %v\n", bytes)
}
