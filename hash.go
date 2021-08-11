package iutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - (len(ciphertext) % blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}


func AesCbcEncrypt(str string, key []byte, IV []byte) ([]byte, error) {
	origData := []byte(str)

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := cipherBlock.BlockSize()
	origData = PKCS5Padding(origData, blockSize)

	crypted := make([]byte, len(origData))
	cipher.NewCBCEncrypter(cipherBlock, IV).CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesCbcDecrypt(encrypted []byte, key []byte, IV []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypt := make([]byte, len(encrypted))

	cipher.NewCBCDecrypter(cipherBlock, IV).CryptBlocks(decrypt, encrypted)
	return decrypt, nil
}
