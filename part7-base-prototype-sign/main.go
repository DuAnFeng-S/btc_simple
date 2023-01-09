package main

import (
	"GoLand_Workspace/part6-base-prototype-wallet/wallet"
	"fmt"
)

/*
公钥私钥
*/
func main() {
	//cmd := BLC.CommandLine{}
	//defer os.Exit(0)
	//cmd.Run()

	//hash := sha256.New()
	//hash.Write([]byte("今天天气不错"))
	//sum := hash.Sum(nil)
	//fmt.Println("hash256:64位16进制字符串", hex.EncodeToString(sum))
	//fmt.Println("hash256:64位16进制字符串", sum)
	//fmt.Printf("hash256:64位16进制字符串 %x\n", sum)

	//hasher := ripemd160.New()
	//hasher.Write([]byte("今天天气不错"))
	//bytes := hasher.Sum(nil)
	//fmt.Printf("ripemd160，40位十六进制字符串,160位二进制,20字节：%x\n", bytes)
	//
	//tt := sha256.New()
	//tt.Write(bytes)
	//i := tt.Sum(nil)
	//
	//jj := sha256.New()
	//jj.Write(i)
	//i2 := jj.Sum(nil)
	//fmt.Println(i2)
	//fmt.Printf("%x\n", i2)

	// base58  64 对称加密，都需要一个种子语句
	newWallet := wallet.NewWallet()
	key := newWallet.PrivateKey
	publicKey := newWallet.PublicKey
	fmt.Println("私钥是", key)
	fmt.Printf("公钥是：%x\n", publicKey)
	address := newWallet.GetAddressfromPublic()
	fmt.Println("钱包的地址是：", address)
}
