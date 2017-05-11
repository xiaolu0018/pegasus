package sms

import (
	"crypto/aes"
	"crypto/cipher"
)

type AesEncrypt struct {
}

//todo strKey 应该可配置
func (a *AesEncrypt) getKey() []byte {
	strKey := "YJ3PeDkKBz4rY6mS"
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//通过CBC/NOPadding 加密 加密字符串

func (a *AesEncrypt) EncryptByCBCNoPadding(plantText []byte) ([]byte, error) {
	key := a.getKey()
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}

	blockModel := cipher.NewCBCEncrypter(block, key)
	plantText = NoPadding(plantText, block.BlockSize())
	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func NoPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := make([]byte, padding)
	return append(ciphertext, padtext...)
}
