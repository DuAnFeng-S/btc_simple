package BLC

// 这种记录的方式无法持久化的保存数据。区块链系统进行调试时每次都需要重新创建区块链，区块链并没有得到保存，这与实际的区块链系统不符
type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) *Block {
	newBlock := NewBlock(bc.Blocks[len(bc.Blocks)-1].Hash, data, bc.Blocks[len(bc.Blocks)-1].Height+1)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

func CreateBlockChain() *BlockChain {
	blockchain := BlockChain{}
	blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	return &blockchain
}
