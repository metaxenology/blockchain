package blockchain

import (
	"crypto/sha256"
	"time"

	"github.com/metaxenology/blockchain/utils"
)

type Block struct {
	timestamp      []byte
	nonce          int64
	data           []byte
	prevHash, hash []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (myBlock *Block) GenerateHash() {
	combinedData := utils.MultiSliceAppend([][]byte{
		myBlock.timestamp,
		utils.ToByteSlice(myBlock.nonce),
		myBlock.data,
		myBlock.prevHash,
	})
	hashResult := sha256.Sum256(combinedData)
	myBlock.hash = hashResult[:]
}

func CreateGenesisBlock() *Block {
	genesisBlock := Block{
		timestamp: utils.ToByteSlice(time.Now().Unix()),
		nonce:     0,
		data:      []byte("Genesis"),
		prevHash:  []byte{},
		hash:      []byte{},
	}
	genesisBlock.GenerateHash()
	return &genesisBlock
}
func InitialiseBlockChain() *BlockChain {
	gen := CreateGenesisBlock()
	myBlockChain := &BlockChain{Blocks: []*Block{gen}}
	return myBlockChain
}
