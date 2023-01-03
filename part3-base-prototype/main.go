package main

import (
	"GoLand_Workspace/part3-base-prototype/BLC"
	"os"
)

/*
把区块记录到数据库中，并同时伴随着我们区块链的组件越来越多，我们需要一个统一的功能管理模块来操作区块链，而不是手动地去调用一个又一个的函数。flag

增加了命令行运行方式和数据持久化到数据库
*/
func main() {
	//con := CommandLine{}

	cmd := BLC.CommandLine{}
	defer os.Exit(0)
	//cmd := CommandLine{}
	cmd.Run()

	// ------------------------------------------test

	//data := "0, 0, 13"
	//fmt.Println([]byte(data))

	// ------------------------------------------test--

	//BLC.Block{}
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	// 打开数据库
	//println("开始创建数据库....")
	//chain := BLC.CreateBlockChain()
	//time.Sleep(6 * time.Second)
	//chain.AddBlock("今天天气不错哟1")
	//time.Sleep(6 * time.Second)
	//chain.AddBlock("今天天气可以2")
	//time.Sleep(6 * time.Second)
	//chain.AddBlock("今天天气可以3")
	//time.Sleep(6 * time.Second)
	//chain.ViewChainData()

	//chain := BLC.ReturnChain()
	//chain.SingleCheck([]byte{0, 0, 3, 28, 43, 200, 25, 102, 91, 121, 246, 23, 228, 36, 193, 142, 8, 188, 246, 51, 210, 117, 77, 136, 239, 105, 79, 153, 28, 216, 38, 47})
	//chain.DBStop()

	//time.Sleep(5 * time.Second)
	//println("开始读取数据。。。。")
	//db, err := bolt.Open(BLC.ChainName, 0600, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//

	/*单节点检查*/
	//db, err := bolt.Open(BLC.ChainName, 0600, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = db.View(func(tx *bolt.Tx) error {
	//	println("尝试进入表")
	//	table := tx.Bucket([]byte(BLC.TableName))
	//
	//	if table != nil {
	//		println("尝试通过key拿值")
	//		v := table.Get([]byte{0, 0, 10, 194, 110, 23, 249, 81, 204, 66, 155, 104, 18, 162, 245, 181, 104, 247, 242, 41, 234, 169, 86, 227, 154, 242, 39, 140, 216, 11, 0, 88})
	//		fmt.Println("进行反序列化...")
	//		block := BLC.DeSerializeBlock(v)
	//		fmt.Println("反序列化后的block为：", block)
	//		fmt.Println("data内容是", string(block.Data))
	//		fmt.Println("区块高度为：", block.Height)
	//		fmt.Println("nonce值是：", block.Nonce)
	//		fmt.Println("前节点是：", []byte(block.PreHash))
	//		//fmt.Println(block.Hash)
	//		//fmt.Println(block)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err)
	//}

	//创建表
	//err1 := db.Update(func(tx *bolt.Tx) error {
	//	//打开表
	//	b := tx.Bucket([]byte("MyBucket"))
	//	//往表里面存储数据
	//	if b != nil {
	//		err := b.Put([]byte("db"), []byte("42"))
	//		if err != nil {
	//			log.Panic("数据存储失败", err)
	//		}
	//	}
	//	return nil
	//})

}
