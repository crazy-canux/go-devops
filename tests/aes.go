package tests

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	plaintext := []byte("my name is canux")
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}
	key_text := "canuxtestijajiasg1234123"
	if len(os.Args) > 2 {
		key_text = os.Args[1]
	}
	fmt.Println(len(key_text))

	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewChiper(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}