package BLC

import (
	"time"
)

// byte打印的是ascii码
type Block struct {
	Height    int64
	preHash   []byte
	Timestamp int64
	Data      []byte
	Hash      []byte
	Nonce     int64 //This line is new
	//Target    int64
}

func NewBlock(preHash []byte, data string, height int64) *Block {

	block := &Block{height, preHash, time.Now().Unix(), []byte(data), []byte{}, 0}
	//block.SetHash()
	//println("区块高度：", block.Height)
	//println("区块前节点hash：", block.preHash)
	//println("区块时间戳：", block.Timestamp)
	//println("区块数据字节：", block.Data)
	//println("区块运算后的hash字节：", block.Hash)
	//println("data还原后：", string([]byte(data)))
	block.Nonce = block.mine()
	block.SetHash()
	//fmt.Println(block)
	return block
}

func GenesisBlock() *Block {
	genesisWords := "Hello, blockchain!"
	return NewBlock([]byte{}, genesisWords, 0)
}
