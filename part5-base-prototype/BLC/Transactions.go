package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type Transaction struct {
	Hash    []byte
	Inputs  []*TxInput
	Outputs []*TxOutput
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
	tx := Transaction{[]byte("这是创世区块的交易"), []*TxInput{txIn}, []*TxOutput{txOut}}
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
func (bc *BlockChain) CreateTransaction(from, to string, amount int) (Transaction, bool) {

	//通过一个方法，返回总余额和一个字典，字典中包括的是一个交易hash和对应的未花费的索引值。总余额是根据所有为花费的交易的总和

	/*
		1.找到from者的所有未花费的交易和余额，如果金额不够，直接return nil ,false
		2.从前往后遍历，找到所有能凑到amount的txAmount
		3.设置tx
		// &TxInput{[]byte{0}, index, []byte(from)}
		// &TxOutput{amount, []byte(to)}
		// &TxOutput{txAmount - amount, []byte(from)}
	*/
	//var txn *Transaction
	var txInputs []*TxInput
	var txOutputs []*TxOutput
	outAmount, outMap := bc.FindSpendableOutputs(from, amount)
	println("找到的金额是：", outAmount)

	if outAmount < amount {
		fmt.Println("Not enough coins!")
		return Transaction{}, false
	}

	for txid, outidx := range outMap {
		//转换为16进制
		//txID, err := hex.DecodeString(txid)
		//Handle(err)
		input := &TxInput{[]byte(txid), outidx, []byte(from)}
		txInputs = append(txInputs, input)
	}

	txOutputs = append(txOutputs, &TxOutput{amount, []byte(to)})
	if outAmount > amount {
		txOutputs = append(txOutputs, &TxOutput{outAmount - amount, []byte(from)})
	}
	println("新的Amount", outAmount-amount)

	txn := Transaction{nil, txInputs, txOutputs}
	txn.SetTransactionHash()

	//println("创建的交易的样子：")
	//println(txn.Hash)
	//println(txn.Inputs)
	//println(txn.Outputs)
	//println("创建一笔交易成功")
	return txn, true
}
