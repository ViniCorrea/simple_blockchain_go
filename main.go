package main

import "fmt"

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Add a block to the blockchain
	blockchain.AddBlock("Send 1 Coin to Fulano")
	blockchain.AddBlock("Send 1 Coin to Fulano")
	blockchain.AddBlock("Send 5 Coin to Sicrano")
	blockchain.AddBlock("Send 1 Coin to Fulano")

	// Print the blocks of the blockchain
	for _, block := range blockchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
