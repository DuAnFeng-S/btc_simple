package main

import (
	"GoLand_Workspace/part4-base-prototype/BLC"
	"os"
)

/*
实现单次输入的交易信息进行发送功能
FindSpendableOutputs:找到最小的符合输出的map
SendTransactions
完成查询和金额不够的转换
** 注意[]byte转换为十六进制位字符的时候，先转换为string

*/
func main() {
	cmd := BLC.CommandLine{}
	defer os.Exit(0)
	//cmd := CommandLine{}
	cmd.Run()
	//var text string = `"[1,2,3,4]"`
	//BLC.JsonToArray(text)
}
