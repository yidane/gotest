package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"time"
)

/*
HELP!!!
http://blog.csdn.net/u012104691/article/details/53766045
*/

var CommentURL string = ""

func AESDecrypt(origData, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func createKey() string {
	const keyArray string = "abcdefghijklmnopqrstuvwxyz0123456789"
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	buf := bytes.Buffer{}
	for i := 0; i < 16; i++ {
		ri := rand.Intn(35)
		buf.WriteString(keyArray[ri : ri+1])
	}

	return buf.String()
}
