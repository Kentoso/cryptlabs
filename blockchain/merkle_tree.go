package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func CalculateMerkleRoot(transactions []Transaction) []byte {
	if len(transactions) == 0 {
		return []byte{}
	}

	var txHashes [][]byte
	for _, tx := range transactions {
		hash := sha256.Sum256([]byte(fmt.Sprintf("%s%s%d%s", tx.Sender, tx.Receiver, tx.Amount, tx.Timestamp.UTC().Format(time.RFC3339))))
		txHashes = append(txHashes, hash[:])
	}

	for len(txHashes) > 1 {
		var newLevel [][]byte
		for i := 0; i < len(txHashes); i += 2 {
			if i+1 == len(txHashes) {
				hash := sha256.Sum256(append(txHashes[i], txHashes[i]...))
				newLevel = append(newLevel, hash[:])
			} else {
				hash := sha256.Sum256(append(txHashes[i], txHashes[i+1]...))
				newLevel = append(newLevel, hash[:])
			}
		}
		txHashes = newLevel
	}
	return txHashes[0]
}
