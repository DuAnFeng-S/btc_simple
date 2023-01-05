package main

import (
	"GoLand_Workspace/part3-base-prototype/BLC"
	"os"
)

/*
transaction：在一个区块中放入一个transaction数组，
1.添加Transaction
2.实现GetBalance，利用UnSpentTx找到address所有未消费的tx，进行统计
** 在控制台输入到程序时，flag包自动将引号剥去，导致字符串转数组失败
** UTXO的通过为花费的金额的时候完美的利用的数组索引
*/
func main() {
	cmd := BLC.CommandLine{}
	defer os.Exit(0)
	//cmd := CommandLine{}
	cmd.Run()
	//var text string = `"[1,2,3,4]"`
	//BLC.JsonToArray(text)
}
