package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
	"strconv"
	"time"
)

// 这种记录的方式无法持久化的保存数据。区块链系统进行调试时每次都需要重新创建区块链，区块链并没有得到保存，这与实际的区块链系统不符
type BlockChain struct {
	//Blocks   []*Block
	Tip      []byte
	DataBase bolt.DB
}

// 把创建的交易数组传入进行挖矿
//
func (bc *BlockChain) Mine(txs []*Transaction) {
	/*
		自己去交易池中拿到交易池中的数据，根据自己的条件自己去交易池中添加交易到本次的数组
	*/
	//var txs []*Transaction

	//println("开始写交易到数据库")
	bc.SendTransactions(txs)
}

// 创建一个交易池

// 创建交易数组 由cli的外部调用这个方法，解析所有的交易数据，
func (bc *BlockChain) CreateTransactions(from, to, amount []string) []*Transaction {

	txs := []*Transaction{}
	// 取到每一组的form，to，amount，把它生成tx。然后把他追加到txs中
	for i := 0; i < len(from); i++ {
		//println("查找所有未花费的交易")
		unSpentTx := bc.UnSpentTx(from[i])
		fmt.Println(unSpentTx)
		// 把数组打包成交易
		amm, err := strconv.Atoi(amount[i])
		Handle(err)
		//println("开始创建每一笔交易")
		//println(from[i])
		//println(to[i])
		//println(amm)
		tx, flag := bc.CreateTransaction(from[i], to[i], amm)
		if flag == false {
			continue
		}
		println("交易生成结果", flag)
		txs = append(txs, &tx)
	}
	return txs

	// 把交易发送到数据库
	//bc.SendTransactions(txs)
}

// 发送交易到链
func (bc *BlockChain) SendTransactions(txns []*Transaction) {

	var newBlock *Block
	//blockchain := BlockChain{}
	err := bc.DataBase.Update(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(TableName))

		if table != nil {

			v := table.Get(bc.Tip)
			//println("2", BytesToInt(bc.Tip))
			preBlock := DeSerializeBlock(v)
			//println("上一个区块的高度:", preBlock.Height)
			//println("3")
			//println(string(bc.Tip[:]))
			fmt.Println(bc.Tip)
			newBlock = NewBlock(bc.Tip, txns, preBlock.Height+1)
			//println(newBlock.Height)
			//if newBlock != nil {
			//	fmt.Println("区块创建成功")
			//} else {
			//	fmt.Println("区块链创建失败")
			//}
			//fmt.Println("写入数据库的block数据为：", newBlock)
			table.Put(newBlock.Hash, newBlock.Serialize())
			bc.Tip = newBlock.Hash

			err := table.Put([]byte(LastHash), newBlock.Hash)
			Handle(err)
		}
		return nil
	})
	Handle(err)
	//return &blockchain
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

			fmt.Println("交易Input:")
			for j := 0; j < len(block.Transactions[i].Inputs); j++ {
				fmt.Println("----------------------------区块", d, "中的第", i, "笔交易的第", j, "个output-------------------------------------")
				fmt.Println("fromAddress:", string(block.Transactions[i].Inputs[j].People[:]))
				fmt.Println("消耗的交易hash:", hex.EncodeToString(block.Transactions[i].Inputs[j].Hash))
				fmt.Println("在一笔交易中TxOutput的索引值:", block.Transactions[i].Inputs[j].Index)
			}

			fmt.Println("")
			fmt.Println("交易Output:")
			for k := 0; k < len(block.Transactions[i].Outputs); k++ {
				fmt.Println("----------------------------区块", d, "中的第", i, "笔交易的第", k, "个output-------------------------------------")
				fmt.Println("Amount:", block.Transactions[i].Outputs[k].ReceiveAmount)
				fmt.Println("toAddress:", string(block.Transactions[i].Outputs[k].People[:]))
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
				通过两层遍历，拿到这笔交易的output数组和每一笔交易hash
				再循环数组，得到每一个output的。通过查看spendTxs[区块hash]来判断index是否和output的索引相等，如果相等则说明被消费，进入下一次循环，否则：
				判断这笔交易的receiver是不是等于address，等于则添加到unSpent。
				最后再另一个for中得到所有的inputs，如果sender是address，则把数据添加到spentTxs[in.hash] = in.outindex
			spendTxs的数据哪里来的？具体过程思考
		1. 第一层for循环：得到每一个block
		2. 第二次循环：将每一个block中的Txs数组循环。并拿到本笔tx的Hash
		3.lable 第三层循环：将tx的Outputs数组循环，如果在spentTxs中存在，则在通过spentTxs[txHash]循环看是否能找到和outIndex相等的值，如果存在则跳过本次循环
		4.否则判断地址是否相等，相等的话将本笔tx加入upSpentTxs
		5.第二层： 查看本笔交易是不是创世交易，如果不是才可以循环inputs，然后把交易的inputs中sender等于address的人加入spentTxs
		6.第一层：当迭代到第一层的时候退出
		7.函数中返回所有的unSpentTxs
	*/
	var unSpentTxs []*Transaction
	spentTxs := make(map[string][]int) // 存储的是每个tx中，已经被消费的output数组的数组序列
	iterator := bc.initChainIterator()
all:
	for {
		block := iterator.Iterator()
		for _, tx := range block.Transactions {
			txHash := hex.EncodeToString(tx.Hash)
		IterOutputs:
			for outIndex, people := range tx.Outputs {
				fmt.Println("进入了out的交易循环")
				if spentTxs[txHash] != nil {
					for _, spentOut := range spentTxs[txHash] {
						if spentOut == outIndex {
							continue IterOutputs
						}
					}
				}
				if people.ToAddressRight([]byte(address)) {
					fmt.Println("找到了一笔交易并成功添加到数组")
					unSpentTxs = append(unSpentTxs, tx)
				}
			}

			if !tx.IsBase() {
				for _, in := range tx.Inputs {
					if in.FromAddressRight([]byte(address)) {
						inHash := hex.EncodeToString(in.Hash)
						spentTxs[inHash] = append(spentTxs[inHash], in.Index)
					}
				}
			}
		}
		//当迭代到创世区块的时候退出
		var hashInt big.Int
		hashInt.SetBytes(block.PreHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			//fmt.Println("遍历完成...")
			break all
		}
		//if bytes.Equal(block.PreHash, bc.BackOgPrevHash()) {
		//	break all
		//}
	}
	//println(spentTxs)
	return unSpentTxs
}

func (chain *BlockChain) BackOgPrevHash() []byte {
	var ogprevhash []byte
	err := chain.DataBase.View(func(txn *bolt.Tx) error {

		table := txn.Bucket([]byte(TableName))
		item := table.Get([]byte(LastHash))
		//Handle(err)

		ogprevhash = item
		return nil
	})
	Handle(err)

	return ogprevhash
}

//返回所有账户的总余额和未花费的tx
func (bc *BlockChain) GetBalance(address string) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	var sum int
	//拿到所有未花费的txns
	unSpentTxs := bc.UnSpentTx(address)

	//将所有的tx求和
	for _, tx := range unSpentTxs {
		hexT := hex.EncodeToString(tx.Hash)
		for outIndex, out := range tx.Outputs {
			if out.ToAddressRight([]byte(address)) {
				sum += out.ReceiveAmount
				unspentOuts[hexT] = outIndex
			}
		}
	}
	//fmt.Println(unSpentTxs)
	return sum, unspentOuts
}

func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.UnSpentTx(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.Hash)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight([]byte(address)) && accumulated < amount {
				accumulated += out.ReceiveAmount
				unspentOuts[txID] = outIdx
				if accumulated >= amount {
					break Work
				}
				continue Work // one transaction can only have one output referred to adderss
			}
		}
	}
	return accumulated, unspentOuts
}
