package cryptographic

import (
	"testing"
)

var src = "/home/canux/Src/go/src/github.com/go-devops/data/cryptographic/APT.yar"
var dest = "/home/canux/Src/go/src/github.com/go-devops/data/cryptographic/APT.yar.aes"
var key = "yNFX2EerZneJfG2o"

func TestAesEncryptFile(t *testing.T) {
    // 16, 24, 32

	err := EncryptFile(src, dest, key)
	if err != nil {
		t.Error("Encrypt file failed.")
	}
}

func TestAesDecryptFile(t *testing.T) {
	plaintext, err := DecryptFile(dest, key)
	if err != nil {
		t.Error("Decrypt file failed.")
	} else {
		t.Log(plaintext)
	}
}


