package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
)

// 返回一个32字节的hash
// 在我们进行hash运算的时候，是先把所有的数据转换为16进制的字节码
func (b *Block) SetHash() {
	information := bytes.Join([][]byte{ToHexInt(b.Height), b.PreHash, ToHexInt(b.Timestamp), b.Data, ToHexInt(b.Nonce)}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func ToHexInt(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	Handle(err)
	//println("将int64转换为字节串类型：", buff.Bytes())
	return buff.Bytes()
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//block.go
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()
}

func DeSerializeBlock(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block
}
