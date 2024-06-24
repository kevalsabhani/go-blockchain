package main

import (
	"fmt"

	"github.com/kevalsabhani/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("First block after genesis")
	chain.AddBlock("Second block after genesis")
	chain.AddBlock("Third block after genesis")
	chain.AddBlock("Fourth block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("**********************************************************\n")
	}
}
