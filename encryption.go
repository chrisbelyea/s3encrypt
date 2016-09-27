// Package encryption contains all the ECB and CBC encryption routines
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

// BlockSize Export this value (which is always 16 lol) to other packages so they don't need
// to import crypto/aes
var blockSize = aes.BlockSize

// ECBDecrypt This function does the ECB decryption of the stored data encryption key
// with the KMS generated envelope key
func ecbDecrypt(ciphertext []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		genError(errors.New("ECBEncrypt - There was a cipher initialization error"))
	}
	bs := aes.BlockSize
	if len(ciphertext)%bs != 0 {
		genError(errors.New("ECBDecrypt - ciphertext is not multiple of blocksize"))
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
	finalplaintextunpad := unpad(finalplaintext)
	return finalplaintextunpad
}

// ECBEncrypt This function encrypts the data encryption key with the
// KMS generated envelope key
func ecbEncrypt(plaintext []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		genError(errors.New("ECBEncrypt - There was a cipher initialization error"))
	}
	bs := aes.BlockSize
	i := 0
	paddedPlain := pad(plaintext)
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

// DecryptFile This function uses the decrypted data encryption key and the
// retrived IV from the S3 metadata to decrypt the data file
func decryptFile(data []byte, iv []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		genError(errors.New("DecryptFile - There was a cipher initialization error"))
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return unpad(data)
}

// EncryptFile This function uses the provided data encryption key and generates
// an IV to encrypt the data file
func encryptFile(data []byte, key []byte) ([]byte, []byte) {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		genError(errors.New("Encryptfile - There was an IV generation error"))
	}
	pmessage := pad(data)
	ciphertext := make([]byte, len(pmessage))
	c, kerr := aes.NewCipher(key)
	if kerr != nil {
		genError(errors.New("EncryptFile - There was a cipher initialization error"))
	}
	mode := cipher.NewCBCEncrypter(c, iv)
	mode.CryptBlocks(ciphertext, pmessage)
	return ciphertext, iv
}

// generatedatakey Does what's on the tin, generates the data encryption key
func generateDataKey() []byte {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		genError(errors.New("GenerateDataKey - There was a key generation error"))
	}
	return key
}
