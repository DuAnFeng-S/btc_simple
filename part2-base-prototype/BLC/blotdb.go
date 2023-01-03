package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type BlockChainIterator struct {
	CurrentHash []byte
	DataBase    bolt.DB
	//adf      Block
}

func (iterator *BlockChainIterator) Iterator() *Block {
	var block *Block

	err := iterator.DataBase.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TableName))
		if b != nil {
			v := b.Get(iterator.CurrentHash)
			block = DeSerializeBlock(v)
		}
		return nil
	})
	Handle(err)
	iterator.CurrentHash = block.PreHash

	//iterator.DataBase.View()

	return block

}

func (chain *BlockChain) SingleCheck(key []byte) {
	var block *Block
	db, err := bolt.Open(ChainName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.View(func(tx *bolt.Tx) error {
		println("尝试进入表")
		table := tx.Bucket([]byte(TableName))

		if table != nil {
			//println("尝试通过key拿值")
			v := table.Get(key)
			//println("进行反序列化", v)
			block = DeSerializeBlock(v)
			fmt.Println("-----------------------------hash对应区块信息------------------------------------")
			//fmt.Println("迭代器得到的block查看是否有preHash:", block)
			fmt.Printf("data内容是:%s \n", string(block.Data))
			fmt.Printf("区块高度为：%d \n", block.Height)
			fmt.Printf("nonce值是：%d \n", block.Nonce)
			fmt.Println("上一个节点是:", block.PreHash)
			fmt.Println("时间戳是：", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
			fmt.Println("本节点hash:", block.Hash)
			fmt.Println("")
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	//return block
}

func (blockchain *BlockChain) DBStop() {
	err := blockchain.DataBase.Close()
	Handle(err)
	fmt.Println("数据库关闭成功。")
	//defer db.Close()
}

func ReturnChain() *BlockChain {

	var lastHash []byte

	db, err := bolt.Open(ChainName, 0600, nil)
	Handle(err)

	err = db.View(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(TableName))
		//table := tx.Bucket([]byte(LastHash))
		lastHash = table.Get([]byte(LastHash))
		return nil
	})
	Handle(err)

	chain := BlockChain{lastHash, *db}
	fmt.Println("chain的lasthash", chain.Tip)
	return &chain
}
