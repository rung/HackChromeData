package decrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func aes128CBCDecrypt(key, iv, encryptPass []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encryptLen := len(encryptPass)
	if encryptLen < block.BlockSize() {
		return nil, errors.New("length of encrypted password less than block size")
	}

	dst := make([]byte, encryptLen)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, encryptPass)
	dst = pkcs5UnPadding(dst, block.BlockSize())
	return dst, nil
}

func pkcs5UnPadding(src []byte, blockSize int) []byte {
	n := len(src)
	paddingNum := int(src[n-1])
	if n < paddingNum || paddingNum > blockSize {
		return src
	}
	return src[:n-paddingNum]
}
