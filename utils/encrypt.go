package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	// b64 "encoding/base64"
	"encoding/hex"
)

func EncryptDatas(res string) string {
	plaintext := res
	key := []byte("YOUR KEY")
	iv := []byte("WITH IV")
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5PaddingNew([]byte(plaintext), aes.BlockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}
func PKCS5PaddingNew(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func DecryptDatas(encryptedURL string) (string, error) {
	ciphertext, err := hex.DecodeString(encryptedURL)
	if err != nil {
		return "", err
	}

	key := []byte("YOUR KEY")
	iv := []byte("WITH IV")
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5Unpadding(plaintext)

	return string(plaintext), nil
}

func PKCS5Unpadding(plaintext []byte) []byte {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	return plaintext[:(length - unpadding)]
}
