package tests

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"sync"
	"log"
	"os"
	"io/ioutil"
	"flag"
)

var (
	// 16, 24, 32
	key = []byte("yNFX2EerZneJfG2o")
	syncMutex sync.Mutex
)


func main() {
	cipher := flag.String("cipher", "yNFX2EerZneJfG2o", "Specify the cipher")
	source := flag.String("src", "/opt/capture/signature/2.4", "Specify source files.")
	dest := flag.String("dest", "/var/www/capture/signature/2.4", "Speicfy encrypt file destination")

	flag.Parse()

	ioutil.WriteFile()

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	plaintext, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	enString, err := Encrypt(string(plaintext))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(enString)

	enByte, err := hex.DecodeString(enString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(enByte)
	err = ioutil.WriteFile(os.Args[2], enByte, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	ciphertext := make([]byte, aes.BlockSize + len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return hex.EncodeToString(ciphertext), nil
}