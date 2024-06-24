package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
