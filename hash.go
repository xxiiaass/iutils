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

// 计算md5值
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 计算base64
func Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 计算Sha256
func Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - (len(ciphertext) % blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}


// 使用aes-cbc加密方法加密数据
func AesCbcEncrypt(str string, key []byte, IV []byte) ([]byte, error) {
	origData := []byte(str)

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := cipherBlock.BlockSize()
	origData = _PKCS5Padding(origData, blockSize)

	crypted := make([]byte, len(origData))
	cipher.NewCBCEncrypter(cipherBlock, IV).CryptBlocks(crypted, origData)
	return crypted, nil
}
// 使用aes-cbc加密方法解密数据

func AesCbcDecrypt(encrypted []byte, key []byte, IV []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypt := make([]byte, len(encrypted))

	cipher.NewCBCDecrypter(cipherBlock, IV).CryptBlocks(decrypt, encrypted)
	return decrypt, nil
}
