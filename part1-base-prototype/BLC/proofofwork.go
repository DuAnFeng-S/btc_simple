package BLC

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

/*
dev：获取难度值进行左位移后的数值
Lsh函数就是向左移位，移的越多目标难度值越大，哈希取值落在的空间就更多就越容易找到符合条件的nonce。
*/
func (b *Block) GetTarget() []byte {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return target.Bytes()
}

/*
获取拼接的加上nonce值的blcok信息的字节码
*/
func (b *Block) GetBase4Nonce(nonce int64) []byte {
	data :=
		bytes.Join([][]byte{ToHexInt(b.Height), b.preHash, ToHexInt(b.Timestamp), b.Data, ToHexInt(nonce)},
			[]byte{},
		)
	return data
}

/*
dev:找出合适的nonce值
通过GetBase4Nonce的字节hash后的bigint与难度值的bigint数值进行比较，结果小于难度值就符合
原理：
	32字节 = 256位二进制 = 64个十六进制
	难度值的运算方式是前target位是零：设置难度值为2   8 - 2 = 6 向前位移6位 eg
diff 二进制：0100 0000  == 十进制：64
则
hashtarget 二进制：0011 1111 == 十进制： 63

小于64的都算是符合

当 256 - diff，diff越大，难度越高，因为要符合前N位为零的概率越小


*/
func (b *Block) mine() int64 {
	var intHash big.Int
	var intTarget big.Int
	var hash [32]byte
	var nonce int64
	nonce = 0
	intTarget.SetBytes(b.GetTarget())

	for nonce < math.MaxInt64 {
		data := b.GetBase4Nonce(nonce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(&intTarget) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce
}

func (b *Block) ValidatePoW() bool {
	var intHash big.Int
	var intTarget big.Int
	var hash [32]byte
	intTarget.SetBytes(b.GetTarget())
	data := b.GetBase4Nonce(b.Nonce)
	hash = sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	if intHash.Cmp(&intTarget) == -1 {
		return true
	}
	return false
}
