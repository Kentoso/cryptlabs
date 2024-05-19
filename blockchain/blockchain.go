package blockchain

import (
	"bytes"
	"fmt"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
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

		// Recalculate the hash of the current block and compare with the stored hash
		expectedHash := currentBlock.CalculateHash()
		if !bytes.Equal(currentBlock.Hash, expectedHash) {
			return false, fmt.Sprintf("Invalid hash at block %d: expected %x, got %x", i, expectedHash, currentBlock.Hash)
		}

		// Check if the current block's previous hash matches the previous block's hash
		if !bytes.Equal(currentBlock.PrevHash, previousBlock.Hash) {
			return false, fmt.Sprintf("Invalid previous hash at block %d: expected %x, got %x", i, previousBlock.Hash, currentBlock.PrevHash)
		}
	}
	return true, ""
}
