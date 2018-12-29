package cryptographic

import (
	"testing"
)

var src = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/cryptographic/APT.yar"
var dest = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/cryptographic/APT.yar.aes"
var verify = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/cryptographic/APT1.yar"

// 16, 24, 32
var key = "aNFX2EerZneJfG2z"

func TestAesDecryptFile(t *testing.T) {
	err := DecryptFile(dest, verify, key)
	if err != nil {
		t.Error("Decrypt file failed.")
	}
}


