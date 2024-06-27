package blockchain

type Block struct {
	Hash     []byte
	PrevHash []byte
	Data     []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, prevHash, []byte(data), 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
