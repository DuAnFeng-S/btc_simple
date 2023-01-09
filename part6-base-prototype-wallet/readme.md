## 思维浏览
**生成钱包地址**

1. 私钥，利用椭圆曲线加密库生成ecdsa
2. 公钥，通过私钥生成
3. address：base58(version(2个十六进制字符串)+公钥加密(sha256+160)+Checknumber(sha256+sha256+公钥加密))26字节
4. 