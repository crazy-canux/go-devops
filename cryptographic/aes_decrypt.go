package cryptographic

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
	"errors"
	//"fmt"
	//"encoding/hex"
)

// decrypt file with aes and save to another file.
func DecryptFile(src, dest, key string) error {

	cipherText, err := ioutil.ReadFile(src)
	if err != nil {
		log.Println("Read file failed.")
		return err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println("Create cipher block failed.")
		return err
	}

	if len(cipherText) < aes.BlockSize {
		log.Println("Key must be 16,24,32.")
		return errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(cipherText, cipherText)

	err = ioutil.WriteFile(dest, cipherText, 0755)
	if err != nil {
		log.Println("Save decrypted file failed.")
		return err
	}
	return nil
}
