package BLC

import (
	"encoding/hex"
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

// 挖区块
func (bc *BlockChain) MineNewBlock(from, to, amount []string) {

	txs := []*Transaction{nil}
	// 取到每一组的form，to，amount，把它生成tx。然后把他追加到txs中
	for i := 0; i < len(from); i++ {
		unSpentTx := bc.UnSpentTx(from[i])
		fmt.Println(unSpentTx)
		tx, _ := bc.CreateTransaction([]byte{0}, []byte{0}, 1)
		txs = append(txs, tx)
	}

	// 把数组打包成交易

	// 把交易发送到数据库
	bc.sendToDB(txs)
}

// 发送交易到链
func (bc *BlockChain) sendToDB(txns []*Transaction) {

	err := bc.DataBase.Update(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(TableName))

		if table != nil {
			// 读取最后一个区块的数据
			v := table.Get(bc.Tip)
			preBlock := DeSerializeBlock(v)
			newBlock := NewBlock(bc.Tip, txns, preBlock.Height+1)
			//fmt.Println("写入数据库的block数据为：", newBlock)
			table.Put(newBlock.Hash, newBlock.Serialize())
			bc.Tip = newBlock.Hash
			err := table.Put([]byte(LastHash), newBlock.Hash)
			Handle(err)
		}
		return nil
	})
	Handle(err)
	//defer db.Close()

}

// 创建区块链
func CreateBlockChain(address string) *BlockChain {

	blockchain := BlockChain{}
	//blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	db, err := bolt.Open(ChainName, 0600, nil)
	Handle(err)
	block := GenesisBlock(address)

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

// 初始化迭代器
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
	d := 0
	for {
		block := iterator.Iterator()
		fmt.Println("----------------------------", d, "-------------------------------------")

		//fmt.Println("迭代器得到的block查看是否有preHash:", block)
		fmt.Printf("区块高度为：%d \n", block.Height)
		fmt.Printf("nonce值是：%d \n", block.Nonce)
		fmt.Println("上一个节点是:", hex.EncodeToString(block.PreHash))
		fmt.Println("时间戳是：", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Println("本节点hash:", hex.EncodeToString(block.Hash))
		for i := 0; i < len(block.Transactions); i++ {
			fmt.Println("----------------------------区块", d, "中的第", i, "笔交易-------------------------------------")
			fmt.Println("本交易的hash", hex.EncodeToString(block.Transactions[i].Hash))

			fmt.Println("交易Output:")
			for k := 0; k < len(block.Transactions[k].Inputs); k++ {
				fmt.Println("Amount:", block.Transactions[i].Outputs[k].ReceiveAmount)
				fmt.Println("toAddress:", string(block.Transactions[i].Outputs[k].People[:]))
			}
			fmt.Println("")

			fmt.Println("交易Input:")
			for j := 0; j < len(block.Transactions[j].Inputs); j++ {
				fmt.Println("fromAddress:", string(block.Transactions[i].Inputs[j].People[:]))
				fmt.Println("消耗的交易hash:", hex.EncodeToString(block.Transactions[i].Inputs[j].Hash))
				fmt.Println("在一笔交易中TxOutput的索引值:", block.Transactions[i].Inputs[j].Index)
			}

		}
		fmt.Println("")

		var hashInt big.Int
		hashInt.SetBytes(block.PreHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("遍历完成...")
			break
		}
		d++
	}

}

// 返回所有未花费的TxOutput
func (bc *BlockChain) UnSpentTx(address string) []*Transaction {
	/*
			从后往前遍历，得到所有不在spendTxs中的数据。这样开发的理由：先存在，再使用
		spendTxs的数据哪里来的？具体过程思考

	*/
	var unSpentTxs []*Transaction
	spentTxs := make(map[string][]int) // can't use type []byte as key value
	println(spentTxs)
	return unSpentTxs
}

//返回所有账户的总余额
func (bc *BlockChain) GetBalance(address string) *int {

	unSpentTxs := bc.UnSpentTx(address)
	fmt.Println(unSpentTxs)
	return nil
}
