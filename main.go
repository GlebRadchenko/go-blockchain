package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {
	blockchain := blockchain.NewBlockchain(4)

	blockchain.AddBlock("Alice", "Bob", 1)
	blockchain.AddBlock("Bob", "Alice", 3)
	blockchain.AddBlock("Alice", "Bob", 3)
	blockchain.AddBlock("Bob", "Alice", 7)

	fmt.Println(blockchain.IsValid())

	for _, block := range blockchain.Chain {
		fmt.Println(block.Hash)
	}
}
