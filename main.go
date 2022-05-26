package main

import (
	"fmt"

	"github.com/metaxenology/blockchain/blockchain"
)

func main() {
	myChain := blockchain.InitialiseBlockChain()
	for _, block := range myChain.Blocks {
		fmt.Printf("%+v", *block)
	}
}
