package main

import (
	"GoLand_Workspace/part1-base-prototype/BLC"
	"fmt"
)

/*
完成最简单的区块链的生成工作
*/
func main() {
	//timeString := strconv.FormatInt(time.Now().Unix(), 2)
	//fmt.Println(timeString)
	//fmt.Println([]byte(timeString))
	//fmt.Println("hello world")

	blockchain := BLC.CreateBlockChain()
	//time.Sleep(time.Second)
	block1 := blockchain.AddBlock("block1")
	fmt.Printf("hash: %x\n", block1.Hash)
	//time.Sleep(time.Second)
	block2 := blockchain.AddBlock("block2")
	fmt.Printf("hash: %x\n", block2.Hash)
	//blockchain.AddBlock("block3")
	//验证
	fmt.Println(block1.ValidatePoW())
	fmt.Println(block2.ValidatePoW())

	//block := BLC.NewBlock([]byte{}, "hello go-blockchain", 1)
	//fmt.Println(block)
	//fmt.Println()
	//
	//block1 := BLC.NewBlock(block.Hash, "hello go1", 2)
	//fmt.Println("第二块数据：")
	//fmt.Println(block1)
}
