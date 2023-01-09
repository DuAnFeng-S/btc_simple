package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	Hash    []byte
	Inputs  []*TxInput
	Outputs []*TxOutput
}

type Transactions struct {
	Txs []*Transaction
}

type TxInput struct {
	Hash   []byte
	Index  int
	People []byte
}

type TxOutput struct {
	ReceiveAmount int
	People        []byte
}

//transaction.go
func (tx *Transaction) TxHash() []byte {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	return hash[:]
}

func (tx *Transaction) SetTransactionHash() {
	tx.Hash = tx.TxHash()
}

func BaseTx(toaddress []byte) *Transaction {
	txIn := &TxInput{[]byte{}, -1, []byte{}}
	txOut := &TxOutput{InitCoin, toaddress}
	//toString := hex.EncodeToString([]byte("杜岸峰的区块链"))
	tx := Transaction{[]byte("杜岸峰的区块链"), []*TxInput{txIn}, []*TxOutput{txOut}}
	tx.SetTransactionHash()
	return &tx
}

func (tx *Transaction) IsBase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].Index == -1
}

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(in.People, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(out.People, address)
}

// 根据输入创建交易
//func (txs *Transaction) CreateTransaction(from, to string, amount int) (Transaction, bool) {
func (bc *BlockChain) CreateTransaction(from, to string, amount int, txs []*Transaction) (Transaction, bool) {

	//通过一个方法，返回总余额和一个字典，字典中包括的是一个交易hash和对应的未花费的索引值。总余额是根据所有为花费的交易的总和

	/*
		1.找到from者的所有未花费的交易和余额，如果金额不够，直接return nil ,false
		2.从前往后遍历，找到所有能凑到amount的txAmount
		3.设置tx
		// &TxInput{[]byte{0}, index, []byte(from)}
		// &TxOutput{amount, []byte(to)}
		// &TxOutput{txAmount - amount, []byte(from)}
	*/

	/*
		先从txs中倒叙输出交易，先和txs中的进行比较，然后再进入数据库的tx比较

	*/
	if from == to {
		//println("发送地址和接受地址不能相同")
		return Transaction{}, false
	}
	var txInputs []*TxInput
	var txOutputs []*TxOutput
	outAmount, outMap := bc.FindSpendableOutputs(from, amount, txs)
	//println("对金额的查找完成")

	if outAmount < amount {
		fmt.Println("没有足够的金额！", from)
		return Transaction{}, false
	}

	for txid, outidx := range outMap {

		//decodeString, _ := hex.DecodeString(txid)
		//fmt.Println("hash.string:", decodeString)
		//转换为16进制
		//txID, err := hex.DecodeString(txid)
		//Handle(err)
		//fmt.Println("最小未花费的交易数组：", txID, ":", outidx)

		input := &TxInput{[]byte(txid), outidx, []byte(from)}
		txInputs = append(txInputs, input)
	}

	txOutputs = append(txOutputs, &TxOutput{amount, []byte(to)})
	if outAmount > amount {
		txOutputs = append(txOutputs, &TxOutput{outAmount - amount, []byte(from)})
	}

	txn := Transaction{nil, txInputs, txOutputs}
	//fmt.Println("设置完一笔交易hash")
	txn.SetTransactionHash()

	return txn, true
}

// 数组中，针对这个地址的未花费的交易输出
func GetTxsFromArraytxs(address string, txs []*Transaction) []*Transaction {

	var unSpentTxs []*Transaction

	spentTxs := make(map[string][]int)

	for i := len(txs) - 1; i >= 0; i-- {
		println("进入了循环：", i)
		txHash := txs[i].Hash
		th := hex.EncodeToString(txHash)
	IterOutputs:
		for outIdx, out := range txs[i].Outputs {

			if spentTxs[th] != nil { //失败点

				for _, spentOut := range spentTxs[th] {
					if spentOut == outIdx {
						continue IterOutputs
					}
				}
			}

			if out.ToAddressRight([]byte(address)) {

				unSpentTxs = append(unSpentTxs, txs[i])
			}
		}
		//if !txs[i].IsBase() {
		for _, in := range txs[i].Inputs {
			if in.FromAddressRight([]byte(address)) {
				inTxID := hex.EncodeToString(in.Hash)
				spentTxs[inTxID] = append(spentTxs[inTxID], in.Index)
			}
		}
	}
	return unSpentTxs
}
