package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"os"

	"github.com/Long-Software/Bex/apps/cmd/everia/models"
	"golang.org/x/crypto/pbkdf2"
)

type Vault struct {
	Entries []models.PasswordEntry `json:"entries"`
}

const (
	saltSize   = 16
	keySize    = 32
	iterations = 100_000
)

// Derives AES key from password + salt
func deriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)
}

func EncryptVault(vault Vault, password string) ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	key := deriveKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, data, nil)
	output := append(salt, nonce...)
	output = append(output, ciphertext...)
	return output, nil
}

func DecryptVault(data []byte, password string) (Vault, error) {
	var vault Vault

	if len(data) < saltSize+12 {
		return vault, errors.New("invalid data")
	}

	salt := data[:saltSize]
	nonce := data[saltSize : saltSize+12]
	ciphertext := data[saltSize+12:]

	key := deriveKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return vault, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return vault, err
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return vault, err
	}

	err = json.Unmarshal(plaintext, &vault)
	return vault, err
}

func SaveEncryptedFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}

func LoadEncryptedFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
