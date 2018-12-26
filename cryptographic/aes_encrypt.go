package cryptographic

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io/ioutil"
	"log"
	"io"
)

func EncryptFile(src, dest, key string) error {
	plaintext, err := ioutil.ReadFile(src)
	if err != nil {
		log.Println("Read file failed.")
		return err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println("Create cipher block failed.")
		return err
	}

	cipherText := make([]byte, aes.BlockSize + len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Println("Read rand failed.")
		return err
	}

	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText[aes.BlockSize:], []byte(plaintext))

	err = ioutil.WriteFile(dest, cipherText, 0755)
	if err != nil {
		log.Println("Save encrypted file failed.")
		return err
	}

	return nil
}
