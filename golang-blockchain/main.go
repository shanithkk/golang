package main

import (
	"fmt"
	"strconv"

	"github.com/shanithkk/golang/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Seconds Block after Genesis")
	chain.AddBlock("Third Block ")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash : %x\n", block.PrevHash)
		fmt.Printf("Data in Block : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
