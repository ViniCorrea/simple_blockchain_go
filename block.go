package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Simplified version of blockchain, which contains only significant information.
// Block keeps block headers
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// ProofOfWork
// Take block fields, concatenate them, and calculate a SHA-256 hash on the concatenated combination.
// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	// Get the timestamp as a string
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// Create the headers by joining the PrevBlockHash, Data, and Timestamp
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	// Hash the headers
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Create a NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}

// In any blockchain, there must be at least one block, and such block, the first in the chain
// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
