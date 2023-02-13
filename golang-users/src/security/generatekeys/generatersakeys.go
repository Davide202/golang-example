package generatekeys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	_ "io"
	"log"
	"os"
	_ "strings"

	"golang.org/x/crypto/ssh"
)

//const private_key = "privateKey"

var (
	PRIVATE_KEY *rsa.PrivateKey
	PUBLIC_KEY  *rsa.PublicKey
)

func init() {

	var err error
	private, err := os.ReadFile("./private.pem")
	if err != nil || private == nil {
		generateNewKeyPairs()
	}
	log.Printf("File %T is present", private)
	privatekey := loadRsaPrivateKey()
	if privatekey == nil {
		panic("ERROR READING PRIVATE.PEM FILE")
	}
	PRIVATE_KEY = privatekey
	PUBLIC_KEY = &privatekey.PublicKey

}

func loadRsaPrivateKeyFromBytes(bytes []byte) *rsa.PrivateKey {

	key, err := ssh.ParseRawPrivateKey(bytes)
	if err != nil {
		log.Println("Error Loading Private Key")
		return nil
	}
	return key.(*rsa.PrivateKey)

}

func loadRsaPrivateKey() *rsa.PrivateKey {

	bytes, err := os.ReadFile("./private.pem")
	if err != nil {
		return nil
	}
	key, err := ssh.ParseRawPrivateKey(bytes)
	if err != nil {
		log.Println("Error Loading Private Key")
		return nil
	}
	return key.(*rsa.PrivateKey)

}

func generateNewKeyPairs() {
	log.Println("Generating New Key Pair Files")
	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey
	PRIVATE_KEY = privatekey
	PUBLIC_KEY = publickey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create("private.pem")
	if err != nil {
		log.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		log.Printf("error when encode private pem: %s \n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		log.Printf("error when dumping publickey: %s \n", err)
		os.Exit(1)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create("public.pem")
	if err != nil {
		log.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		log.Printf("error when encode public pem: %s \n", err)
		os.Exit(1)
	}

	log.Println("Generated New Key Pair Files")
}
