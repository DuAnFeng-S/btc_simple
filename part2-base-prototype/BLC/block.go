package BLC

import (
	"fmt"
	"time"
)

// byte打印的是ascii码
type Block struct {
	Height    int64
	PreHash   []byte
	Timestamp int64
	Data      []byte
	Hash      []byte
	Nonce     int64 //This line is new
	//Target    int64
}

func NewBlock(preHash []byte, data string, height int64) *Block {
	//fmt.Println("设置的prehash是:", preHash)

	block := &Block{height, preHash, time.Now().Unix(), []byte(data), []byte{}, 0}
	block.Nonce = block.mine()
	block.SetHash()
	//block.SetHash()
	//fmt.Println("区块高度：", block.Height)
	//fmt.Println("区块前节点hash：", block.PreHash)
	//fmt.Println("区块时间戳：", block.Timestamp)
	//fmt.Println("区块数据字节：", block.Data)
	//fmt.Println("区块运算后的hash字节：", block.Hash)
	//fmt.Println("data还原后：", string([]byte(data)))

	fmt.Println("挖掘区块hash：", block.Hash)
	fmt.Println("创建区块成功")
	return block
}

func GenesisBlock() *Block {
	genesisWords := "Hello, blockchain!"
	block := NewBlock([]byte{}, genesisWords, 0)
	if block != nil {
		fmt.Println("创世区块创建成功")
		return block
	} else {
		fmt.Println("区块链创建失败")
		return nil
	}

}
