package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"go-kms-s3/padding"
)

var BlockSize = aes.BlockSize

func ECB_decrypt(ciphertext []byte, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	bs := aes.BlockSize
	if len(ciphertext)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	i := 0
	plaintext := make([]byte, len(ciphertext))
	finalplaintext := make([]byte, len(ciphertext))
	for len(ciphertext) > 0 {
		cipher.Decrypt(plaintext, ciphertext)
		ciphertext = ciphertext[bs:]
		decryptedBlock := plaintext[:bs]
		for index, element := range decryptedBlock {
			finalplaintext[(i*bs)+index] = element
		}
		i++
		plaintext = plaintext[bs:]
	}
	finalplaintext_unpad := padding.Unpad(finalplaintext)
	return finalplaintext_unpad
}

func ECB_encrypt(plaintext []byte, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	bs := aes.BlockSize
	i := 0
	paddedPlain := padding.Pad(plaintext)
	ciphertext := make([]byte, len(paddedPlain))
	finalciphertext := make([]byte, len(paddedPlain))
	for len(paddedPlain) > 0 {
		cipher.Encrypt(ciphertext, paddedPlain)
		paddedPlain = paddedPlain[bs:]
		encryptedBlock := ciphertext[:bs]
		for index, element := range encryptedBlock {
			finalciphertext[(i*bs)+index] = element
		}
		i++
		ciphertext = ciphertext[bs:]
	}
	return finalciphertext
}

func Decryptfile(data []byte, iv []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return padding.Unpad(data)
}

func Encryptfile(data []byte, key []byte) ([]byte, []byte) {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		panic("There was an IV generation error")
	}
	pmessage := padding.Pad(data)
	ciphertext := make([]byte, len(pmessage))
	c, _ := aes.NewCipher(key)
	mode := cipher.NewCBCEncrypter(c, iv)
	mode.CryptBlocks(ciphertext, pmessage)
	return ciphertext, iv
}
