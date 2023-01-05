package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Transaction struct {
	Hash    []byte
	Inputs  []*TxInput
	Outputs []*TxOutput
}

type TxInput struct {
	Hash   []byte
	Index  int32
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
func (bc *BlockChain) CreateTransaction(from, to []byte, amount int) (*Transaction, bool) {

	//通过一个方法，返回总余额和一个字典，字典中包括的是一个交易hash和对应的未花费的索引值。总余额是根据所有为花费的交易的总和

	//var txn *Transaction
	var txInputs []*TxInput
	var txOutputs []*TxOutput

	txinput := &TxInput{[]byte{0}, 0, from}

	txInputs = append(txInputs, txinput)

	txoutput := &TxOutput{amount, to}
	txOutputs = append(txOutputs, txoutput)

	//txn := &Transaction{[]byte{0}, txInputs, txOutputs}

	return nil, true
}
