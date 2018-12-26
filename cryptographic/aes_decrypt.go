package cryptographic

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
	"errors"
	"fmt"
	//"encoding/hex"
)

func DecryptFile(filename, key string) (string, error) {

	cipherText, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Read file failed.")
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println("Create cipher block failed.")
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		log.Println("Key must be 16,24,32.")
		return "", errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(cipherText, cipherText)
	return fmt.Sprintf("%s", cipherText), nil
}
