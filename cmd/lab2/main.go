package main

import (
	"fmt"
	"github.com/Kentoso/cryptlabs/ops"
)

func alice() {

}

func main() {
	bits := 512
	alicePublic, alicePrivate := ops.RSA(bits, ops.MillerRabin)
	bobPublic, bobPrivate := ops.RSA(bits, ops.MillerRabin)

	// Alice and Bob exchange public keys

	// Alice sends a message to Bob, encoding the message using Bob's public key
	aliceMessage := ops.Encrypt(bobPublic, "Hello, Bob! I'm Alice")

	// Bob receives the message and decodes it using his private key
	decodedAliceMessage := ops.DecryptCRT(bobPrivate, aliceMessage)
	fmt.Println(decodedAliceMessage)

	// Bob sends a message to Alice, encoding the message using Alice's public key
	bobMessage := ops.Encrypt(alicePublic, "Hello, Alice! I'm Bob")

	// Alice receives the message and decodes it using her private key
	decodedBobMessage := ops.DecryptCRT(alicePrivate, bobMessage)
	fmt.Println(decodedBobMessage)

	fmt.Println("Alice's public key:", alicePublic)
	fmt.Println("Alice's private key:", alicePrivate)
	fmt.Println("Bob's public key:", bobPublic)
	fmt.Println("Bob's private key:", bobPrivate)

	//public, private := ops.RSA(512, ops.MillerRabin)
	//message := "Hello, World!"
	//cipher := ops.Encrypt(public, message)
	//decrypted := ops.Decrypt(private, cipher)
	//fmt.Println("Original message:", message)
	//fmt.Println("Encrypted message:", cipher)
	//fmt.Println("Decrypted message:", decrypted)
}
