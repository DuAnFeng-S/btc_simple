package BLC

import (
	"bytes"
	"fmt"
	"time"
)

// byte打印的是ascii码
type Block struct {
	Height    int64
	PreHash   []byte
	Timestamp int64
	//Data      []byte
	Transactions []*Transaction
	Hash         []byte
	Nonce        int64 //This line is new
}

func NewBlock(preHash []byte, txns []*Transaction, height int64) *Block {
	//fmt.Println("设置的prehash是:", preHash)

	block := &Block{height, preHash, time.Now().Unix(), txns, []byte{}, 0}
	block.Nonce = block.mine()
	block.SetHash()
	return block
}

func GenesisBlock(address string) *Block {
	//genesisWords := "Hello, blockchain!"
	tx := BaseTx([]byte(address))
	block := NewBlock([]byte{}, []*Transaction{tx}, 0)
	if block != nil {
		fmt.Println("创世区块创建成功")
		return block
	} else {
		fmt.Println("区块链创建失败")
		return nil
	}

}

// 把交易数组变成字节数组
func (b *Block) BackTrasactionSummary() []byte {
	txIDs := make([][]byte, 0)
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.Hash)
	}
	summary := bytes.Join(txIDs, []byte{})
	return summary
}
