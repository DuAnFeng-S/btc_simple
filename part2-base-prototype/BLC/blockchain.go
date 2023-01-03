package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
	"time"
)

// 这种记录的方式无法持久化的保存数据。区块链系统进行调试时每次都需要重新创建区块链，区块链并没有得到保存，这与实际的区块链系统不符
type BlockChain struct {
	//Blocks   []*Block
	Tip      []byte
	DataBase bolt.DB
}

func (bc *BlockChain) AddBlock(data string) {
	//newBlock := NewBlock(bc.Blocks[len(bc.Blocks)-1].Hash, data, bc.Blocks[len(bc.Blocks)-1].Height+1)
	//bc.Blocks = append(bc.Blocks, newBlock)
	//return newBlock

	//db, err := bolt.Open(ChainName, 0600, nil)
	//Handle(err)

	err := bc.DataBase.Update(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(TableName))

		if table != nil {
			// 读取最后一个区块的数据
			v := table.Get(bc.Tip)
			preBlock := DeSerializeBlock(v)
			newBlock := NewBlock(bc.Tip, data, preBlock.Height+1)
			//fmt.Println("写入数据库的block数据为：", newBlock)
			table.Put(newBlock.Hash, newBlock.Serialize())

			bc.Tip = newBlock.Hash
			//fmt.Println("区块hash：", newBlock.Hash)

			err := table.Put([]byte(LastHash), newBlock.Hash)
			Handle(err)
			//fmt.Println("添加区块成功...")

			//fmt.Println("前一个节点的高度", preBlock.Height)
			//fmt.Println(preBlock.Nonce)
		}
		return nil
	})
	Handle(err)
	//defer db.Close()

}

func CreateBlockChain() *BlockChain {

	blockchain := BlockChain{}
	//blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())

	db, err := bolt.Open(ChainName, 0600, nil)
	Handle(err)
	block := GenesisBlock()

	db.Update(func(tx *bolt.Tx) error {
		table, err2 := tx.CreateBucket([]byte(TableName))
		if err2 != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = table.Put((block.Hash), block.Serialize())
		Handle(err)

		err = table.Put([]byte(LastHash), block.Hash)
		Handle(err)

		return nil
	})

	blockchain.Tip = block.Hash
	blockchain.DataBase = *db

	//defer db.Close()
	return &blockchain
}

func (blockchain *BlockChain) initChainIterator() *BlockChainIterator {

	iterator := BlockChainIterator{blockchain.Tip, blockchain.DataBase}
	return &iterator
}

// 遍历
func (blockchain *BlockChain) ViewChainData() {

	iterator := blockchain.initChainIterator()

	//当hash为0的时候遍历结束
	/*
		1.打印当前节点的信息
		2.用迭代器返回前一个block，Iterator(lasthash)
		3.用前节点继续打印
	*/
	i := 0
	for {
		block := iterator.Iterator()
		fmt.Println("----------------------------", i, "-------------------------------------")
		i++
		//fmt.Println("迭代器得到的block查看是否有preHash:", block)
		fmt.Printf("data内容是:%s \n", string(block.Data))
		fmt.Printf("区块高度为：%d \n", block.Height)
		fmt.Printf("nonce值是：%d \n", block.Nonce)
		fmt.Println("上一个节点是:", block.PreHash)
		fmt.Println("时间戳是：", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Println("本节点hash:", block.Hash)
		fmt.Println("")

		var hashInt big.Int
		hashInt.SetBytes(block.PreHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("遍历完成...")
			break
		}
	}

}
