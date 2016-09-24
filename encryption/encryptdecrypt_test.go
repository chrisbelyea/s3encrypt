package encryption

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestEncryptDecryptCBCOne(t *testing.T) {
	data := []byte("Test String one")
	fmt.Printf("Original Data: %s\n", string(data))
	key := make([]byte, 16)
	rand.Read(key)
	ctext, iv := Encryptfile(data, key)
	result := Decryptfile(ctext, iv, key)
	fmt.Printf("Decrypted Data: %s\n", string(result))
	if string(data) != string(result) {
		t.Error("Decryption failed!  Error in Decryptfile/Encryptfile functions")
	}
}

func TestEncryptDecryptCBCTwo(t *testing.T) {
	data := []byte("abcd1234efgh5678")
	fmt.Printf("Original Data: %s\n", string(data))
	key := make([]byte, 16)
	rand.Read(key)
	ctext, iv := Encryptfile(data, key)
	result := Decryptfile(ctext, iv, key)
	fmt.Printf("Decrypted Data: %s\n", string(result))
	if string(data) != string(result) {
		t.Error("Decryption failed!  Error in Decryptfile/Encryptfile functions")
	}
}
func TestEncryptDecryptECBOne(t *testing.T) {
	data := []byte("Test data string dos")
	fmt.Printf("Original Data: %s\n", string(data))
	key := make([]byte, 16)
	rand.Read(key)
	ctext := ECB_encrypt(data, key)
	result := ECB_decrypt(ctext, key)
	fmt.Printf("Decrypted Data: %s\n", string(result))
	if string(data) != string(result) {
		t.Error("Decryption failed!  Error in ECB_decrypt/ECB_encrypt functions")
	}
}
func TestEncryptDecryptECBTwo(t *testing.T) {
	data := []byte("abcd1234efgh5678")
	fmt.Printf("Original Data: %s\n", string(data))
	key := make([]byte, 16)
	rand.Read(key)
	ctext := ECB_encrypt(data, key)
	result := ECB_decrypt(ctext, key)
	fmt.Printf("Decrypted Data: %s\n", string(result))
	if string(data) != string(result) {
		t.Error("Decryption failed!  Error in ECB_decrypt/ECB_encrypt functions")
	}
}
