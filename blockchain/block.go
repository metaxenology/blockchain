package blockchain

type Block struct {
	timestamp, nonce int64
	data             []byte
	prevHash, hash   []byte
}

type BlockChain struct {
	blocks []*Block
}
