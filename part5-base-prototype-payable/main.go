package main

import (
	"GoLand_Workspace/part5-base-prototype-payable/BLC"
	"os"
)

/*
1.区块交易查找
** 数组中未花费的和数据库中未花费的进行拼接后，要再次进行GetTxsFromArraytxs，这样才能把数组中引用的数据库中的tx删除
** 查找同一个区块的交易的时候，要从后往前找，这样才能接上同一个区块中的已经消费的out

*/
func main() {
	cmd := BLC.CommandLine{}
	defer os.Exit(0)
	cmd.Run()
}
