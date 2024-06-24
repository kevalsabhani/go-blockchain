package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	PrevHash []byte
	Data     []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, prevHash, []byte(data)}
	block.DeriveHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
