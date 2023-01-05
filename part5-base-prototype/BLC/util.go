package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"log"
)

// 返回一个32字节的hash
// 在我们进行hash运算的时候，是先把所有的数据转换为16进制的字节码
func (b *Block) SetHash() {
	information := bytes.Join([][]byte{ToHexInt(b.Height), b.PreHash, ToHexInt(b.Timestamp), b.BackTrasactionSummary(), ToHexInt(b.Nonce)}, []byte{})
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

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//	json转数组
func JsonToArray(jsonString string) []string {
	var sArr []string
	//prefix := bytes.TrimPrefix([]byte(jsonString), []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal([]byte(jsonString), &sArr); err != nil {
		log.Panic(err)
	}
	//fmt.Println(sArr)
	return sArr
}
