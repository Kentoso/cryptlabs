package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Blockchain struct {
	Blocks     []*Block
	MerkleRoot []byte
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{newGenesisBlock()}}
}

func newGenesisBlock() *Block {
	genesisBlock := &Block{
		Transactions: nil,
		Timestamp:    time.Now(),
		PrevHash:     []byte{},
	}
	genesisBlock.MerkleRoot = CalculateMerkleRoot(genesisBlock.Transactions)
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return genesisBlock
}

func (bc *Blockchain) AddBlock(transactions []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Transactions: transactions,
		Timestamp:    time.Now(),
		PrevHash:     prevBlock.Hash,
	}
	newBlock.MerkleRoot = CalculateMerkleRoot(newBlock.Transactions)
	newBlock.MineBlock(4)
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.MerkleRoot = bc.CalculateBlockchainMerkleRoot()
}

func (bc *Blockchain) CalculateBlockchainMerkleRoot() []byte {
	if len(bc.Blocks) == 0 {
		return []byte{}
	}

	var blockHashes [][]byte
	for _, block := range bc.Blocks {
		blockHashes = append(blockHashes, block.Hash)
	}

	for len(blockHashes) > 1 {
		var newLevel [][]byte
		for i := 0; i < len(blockHashes); i += 2 {
			if i+1 == len(blockHashes) {
				hash := sha256.Sum256(append(blockHashes[i], blockHashes[i]...))
				newLevel = append(newLevel, hash[:])
			} else {
				hash := sha256.Sum256(append(blockHashes[i], blockHashes[i+1]...))
				newLevel = append(newLevel, hash[:])
			}
		}
		blockHashes = newLevel
	}

	return blockHashes[0]
}

func (bc *Blockchain) GetAccountsBalances(blockIndex int) (map[string]int, map[string]int, map[string]int) {
	if blockIndex < 0 || blockIndex >= len(bc.Blocks) {
		return nil, nil, nil // Invalid block index
	}

	balances := make(map[string]int)
	minBalances := make(map[string]int)
	maxBalances := make(map[string]int)

	for i := 0; i <= blockIndex; i++ {
		block := bc.Blocks[i]
		for _, tx := range block.Transactions {
			balances[tx.Sender] -= tx.Amount
			balances[tx.Receiver] += tx.Amount

			if currentBalance, exists := minBalances[tx.Sender]; !exists || balances[tx.Sender] < currentBalance {
				minBalances[tx.Sender] = balances[tx.Sender]
			}
			if currentBalance, exists := minBalances[tx.Receiver]; !exists || balances[tx.Receiver] < currentBalance {
				minBalances[tx.Receiver] = balances[tx.Receiver]
			}

			if currentBalance, exists := maxBalances[tx.Sender]; !exists || balances[tx.Sender] > currentBalance {
				maxBalances[tx.Sender] = balances[tx.Sender]
			}
			if currentBalance, exists := maxBalances[tx.Receiver]; !exists || balances[tx.Receiver] > currentBalance {
				maxBalances[tx.Receiver] = balances[tx.Receiver]
			}
		}
	}

	return balances, minBalances, maxBalances
}

func (bc *Blockchain) IsValid() (bool, string) {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		expectedMerkleRoot := CalculateMerkleRoot(currentBlock.Transactions)
		if !bytes.Equal(currentBlock.MerkleRoot, expectedMerkleRoot) {
			return false, fmt.Sprintf("Invalid Merkle root at block %d: expected %x, got %x", i, expectedMerkleRoot, currentBlock.MerkleRoot)
		}

		expectedHash := currentBlock.CalculateHash()
		if !bytes.Equal(currentBlock.Hash, expectedHash) {
			return false, fmt.Sprintf("Invalid hash at block %d: expected %x, got %x", i, expectedHash, currentBlock.Hash)
		}

		if !bytes.Equal(currentBlock.PrevHash, previousBlock.Hash) {
			return false, fmt.Sprintf("Invalid previous hash at block %d: expected %x, got %x", i, previousBlock.Hash, currentBlock.PrevHash)
		}
	}

	expectedBlockchainMerkleRoot := bc.CalculateBlockchainMerkleRoot()
	if !bytes.Equal(bc.MerkleRoot, expectedBlockchainMerkleRoot) {
		return false, fmt.Sprintf("Invalid blockchain Merkle root: expected %x, got %x", expectedBlockchainMerkleRoot, bc.MerkleRoot)
	}

	return true, ""
}
