package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"math/big"
	"time"
)

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    int
	Timestamp time.Time
}

type Block struct {
	Transactions []Transaction
	Timestamp    time.Time
	PrevHash     []byte
	Hash         []byte
	Nonce        int
	MerkleRoot   []byte
}

func (b *Block) CalculateHash() []byte {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(b.Transactions)
	if err != nil {
		panic(err)
	}
	err = enc.Encode(b.Timestamp)
	if err != nil {
		panic(err)
	}
	err = enc.Encode(b.PrevHash)
	if err != nil {
		panic(err)
	}
	err = enc.Encode(b.Nonce)
	if err != nil {
		panic(err)
	}
	data.Write(b.MerkleRoot)

	hash := sha256.Sum256(data.Bytes())
	return hash[:]
}

func (b *Block) MineBlock(difficulty int) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	for {
		hash := b.CalculateHash()
		var hashInt big.Int
		hashInt.SetBytes(hash)
		if hashInt.Cmp(target) == -1 {
			b.Hash = hash
			return
		}
		b.Nonce++
	}
}
