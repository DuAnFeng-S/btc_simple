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
	//println("返回的block：")
	//println("开始挖区块的hash")
	//
	//fmt.Println("本交易的hash", block.Transactions[1].Inputs[0].Index)
	//fmt.Println("本交易的hash", block.Transactions[1].Outputs[0].ReceiveAmount)
	//fmt.Println("本交易的hash", block.Transactions[1].Hash)
	//fmt.Println("区块高度：", block.Height)
	//fmt.Println("区块前节点hash：", block.PreHash)
	//fmt.Println("区块时间戳：", block.Timestamp)
	//fmt.Println("区块运算前的hash字节：", block.Hash)
	fmt.Println("tx的长度", len(block.Transactions))
	//for i := 0; i < len(block.Transactions); i++ {
	//	fmt.Println("----------------------------本区块中的第", i, "笔交易-------------------------------------")
	//	fmt.Println("本交易的hash", block.Transactions[i].Hash)
	//
	//	fmt.Println("交易Output:")
	//	for k := 0; k < len(block.Transactions[k].Inputs); k++ {
	//		fmt.Println("Amount:", block.Transactions[i].Outputs[k].ReceiveAmount)
	//		fmt.Println("toAddress:", string(block.Transactions[i].Outputs[k].People[:]))
	//	}
	//	//fmt.Println("")
	//
	//	fmt.Println("交易Input:")
	//	for j := 0; j < len(block.Transactions[j].Inputs); j++ {
	//		fmt.Println("fromAddress:", string(block.Transactions[i].Inputs[j].People[:]))
	//		fmt.Println("消耗的交易hash:", hex.EncodeToString(block.Transactions[i].Inputs[j].Hash))
	//		fmt.Println("在一笔交易中TxOutput的索引值:", block.Transactions[i].Inputs[j].Index)
	//	}
	//
	//}
	//fmt.Println("data还原后：", string([]byte(data)))
	block.Nonce = block.mine()
	//println("挖区块的hash成功")
	block.SetHash()
	return block
}

func GenesisBlock(address string) *Block {
	//genesisWords := "Hello, blockchain!"
	tx := BaseTx([]byte(address))
	block := NewBlock([]byte{0}, []*Transaction{tx}, 0)
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
