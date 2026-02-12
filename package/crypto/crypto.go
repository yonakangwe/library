package crypto

import (
	"library/package/log"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

/**
 * Generate private and public key using openssl as shown below
 * private key: openssl ecparam -name prime256v1 -genkey -noout -out auth_private_key.pem
 * public key: openssl pkey -in auth_private_key.pem -pubout -out auth_public_key.pem
 * copy these keys into .storage/keys
 * specify the path of these keys into the config.yaml file
 */

// Sign function accepts the message to be signed and the private key and returns the hash and signature
func Sign(message, privKey []byte) (string, string, error) {
	block, _ := pem.Decode(privKey)
	if block == nil {
		log.Error("error decoding client private key")
		return "", "", errors.New("error decoding private key")
	}
	pk, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		log.Errorf("error parsing client private key: %v", err)
		return "", "", err
	}
	// Hash the input to get a summary of the information
	hash := sha256.Sum256([]byte(message))
	// ECDSA Signing
	sign, err := ecdsa.SignASN1(rand.Reader, pk, hash[:])
	if err != nil {
		log.Errorf("error error signing data: %v", err)
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(hash[:]), base64.StdEncoding.EncodeToString(sign), nil
}

func Verify(pubKey []byte, message, sign string) (bool, error) {
	// Load the public key in x509 format
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return false, errors.New("pubKey no pem data found")
	}
	genericPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	pk := genericPublicKey.(*ecdsa.PublicKey)

	// Hash the input to get a summary of the information
	hash, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return false, err
	}
	// ECDSA Validation
	bSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}
	return ecdsa.VerifyASN1(pk, hash[:], bSign), nil
}
