package main

import (
	"fmt"
	"github.com/Kentoso/cryptlabs/ops"
)

func scenario(bits int, print bool) {
	alicePublic, alicePrivate := ops.RSA(bits, ops.MillerRabin)
	bobPublic, bobPrivate := ops.RSA(bits, ops.MillerRabin)

	// Alice and Bob exchange public keys

	// Alice sends a message to Bob, encoding the message using Bob's public key
	aliceMessage := "Hello, Bob! I'm Alice"
	encryptedAliceMessage := ops.Encrypt(bobPublic, aliceMessage)

	// Alice signs the message using her private key
	aliceMessageHash := ops.HashPearson(aliceMessage)
	aliceSignature := ops.Sign(alicePrivate, aliceMessageHash)

	// Bob receives the message and decodes it using his private key
	decodedAliceMessage := ops.DecryptCRT(bobPrivate, encryptedAliceMessage)

	// Bob verifies Alice's signature
	verified := ops.Verify(alicePublic, aliceMessageHash, aliceSignature)

	// Bob sends a message to Alice, encoding the message using Alice's public key
	bobMessage := "Hello, Alice! I'm Bob"
	encryptedBobMessage := ops.Encrypt(alicePublic, bobMessage)

	// Bob signs the message using his private key
	bobMessageHash := ops.HashPearson(bobMessage)
	bobSignature := ops.Sign(bobPrivate, bobMessageHash)

	// Alice receives the message and decodes it using her private key
	decodedBobMessage := ops.DecryptCRT(alicePrivate, encryptedBobMessage)

	// Alice verifies Bob's signature
	verified = ops.Verify(bobPublic, bobMessageHash, bobSignature)

	if print {
		fmt.Println(decodedAliceMessage)
		fmt.Println("Is Alice's message verified?", verified)
		fmt.Println(decodedBobMessage)
		fmt.Println("Is Bob's message verified?", verified)

		fmt.Println("Alice's public key:", alicePublic)
		fmt.Println("Alice's private key:", alicePrivate)
		fmt.Println("Bob's public key:", bobPublic)
		fmt.Println("Bob's private key:", bobPrivate)
	}
}

func main() {
	scenario(512, true)
	//public, private := ops.RSA(512, ops.MillerRabin)
	//message := "Hello, World!"
	//cipher := ops.Encrypt(public, message)
	//decrypted := ops.Decrypt(private, cipher)
	//fmt.Println("Original message:", message)
	//fmt.Println("Encrypted message:", cipher)
	//fmt.Println("Decrypted message:", decrypted)
}
