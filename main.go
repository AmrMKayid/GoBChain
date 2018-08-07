package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1000 BTC to Amr")
	bc.AddBlock("Send more 3000 BTC to Amr :D")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
