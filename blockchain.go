package main

// In its essence blockchain is just a database with certain structure:
// it’s an ordered, back-linked list. Which means that blocks are stored
// in the insertion order and that each block is linked to the previous one.
// This structure allows to quickly get the latest block in a chain and to (efficiently)
// get a block by its hash.

// Just use an array, because we don’t need to get blocks by their hash for now.
type Blockchain struct {
	blocks []*Block
}

// Make it possible to add blocks to it
func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, previousBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Create a NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	// Create a new blockchain with a genesis block
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}