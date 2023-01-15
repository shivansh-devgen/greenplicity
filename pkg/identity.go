package identity

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"github.com/filecoin-project/go-state-types/crypto"
)

// GeneratePrivateKey generates a new private key
func GeneratePrivateKey() (crypto.PrivateKey, error) {
	privateKey := make([]byte, 32)
	_, err := rand.Read(privateKey)
	if err != nil {
		return nil, err
	}
	return crypto.PrivateKey(privateKey), nil
}

// PublicKeyFromPrivateKey returns the public key for a private key
func PublicKeyFromPrivateKey(privateKey crypto.PrivateKey) crypto.PublicKey {
	return crypto.PublicKey(privateKey.ExtractPublic())
}

// SignMessage signs a message with the provided private key
func SignMessage(privateKey crypto.PrivateKey, message []byte) ([]byte, error) {
	return crypto.Sign(privateKey, message)
}

// VerifySignature verifies a signature on a message
func VerifySignature(publicKey crypto.PublicKey, message []byte, signature []byte) error {
	ok, err := crypto.VerifySignature(publicKey, message, signature)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("signature verification failed")
	}
	return nil
}

// HashMessage hashes a message
func HashMessage(message []byte) []byte {
	h := sha256.New()
	h.Write(message)
	return h.Sum(nil)
}

// HashString hashes a string
func HashString(message string) string {
	h := sha256.New()
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

