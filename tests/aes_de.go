package tests

import (
	"crypto/aes"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	//"encoding/hex"
	"crypto/cipher"
	"sync"
)


var (
	key = []byte("iamthekeycanux12")
	syncMutex sync.Mutex
)

flag.S


func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ciphertext, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	deString, err := Decrypt(string(ciphertext))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deString)
}

func Decrypt(de string) (string, error) {
	//cipherText, err := hex.DecodeString(de)
	cipherText := []byte(de)
	//if err != nil {
	//	log.Fatal(err)
	//}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	if len(cipherText) < aes.BlockSize {
		log.Fatal("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	//fmt.Println(len(cipherText), len(iv))
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}