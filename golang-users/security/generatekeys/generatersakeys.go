package generatekeys

import (
	"crypto/rsa"
	_ "io"
	"os"
	"strings"
)

const private_key = "privateKey"

var (
	SignKey   *rsa.PrivateKey
	VerifyKey *rsa.PublicKey
)

func init() {

	if os.Getenv(private_key) == "" {
		os.Setenv(private_key, "sono la chiave privata")
	}

	seed := os.Getenv(private_key)

	reader := strings.NewReader(seed)

	privateKey, err := rsa.GenerateKey(reader, 21)
	if err != nil {
		return
	}
	SignKey = privateKey

	VerifyKey = &SignKey.PublicKey
}
