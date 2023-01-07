## 思维浏览
**发送多笔交易的流程**
cli
1. 传入三个数组并校验
2. 调用CreateTransactions生成交易
3. 调用Mine，对生成的txs进行挖矿

2-1. 遍历出每笔数据调用CreateTransaction生成tx并拼接成txs
2-1-1. 调用 <font color=yellow>FindSpendableOutputs</font> 计算出能满足的未花费的能满足交易需求的最小数组，[txhash] = Outindex
2-1-2. 通过数组,计算出新的交易的信息

2-1-1-1. 计算出所有未花费的交易，包括db和txs中的。两种结果进行拼接。
2-1-1-2. 循环得到能满足amount的数组并返回


