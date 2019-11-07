package cryptographic

import (
	"testing"
)

var src = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/cryptographic/APT.yar"
var dest = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/cryptographic/APT.yar.aes"
var verify = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/cryptographic/APT1.yar"

// 16, 24, 32
var key = "aNFX2EerZneJfG2z"

func TestAesEncryptFile(t *testing.T) {

	err := EncryptFile(src, dest, key)
	if err != nil {
		t.Errorf("Encrypt file failed: %s", err)
	}
}