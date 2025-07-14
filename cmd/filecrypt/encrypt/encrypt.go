package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {
	srcFile, err := os.Open(source)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()
	plain, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}
	key := password
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New())

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	cipher := aesgcm.Seal(nil, nonce, plain, nil)
	cipher = append(cipher, nonce...)

	dstFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	_, err = dstFile.Write(cipher)
	if err != nil {
		panic(err.Error())
	}
}

func Decrypt(source string, password []byte) {
	srcFile, err := os.Open(source)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()
	cipherText, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}
	key := password
	salt := cipherText[len(cipherText)-12:]
	str := hex.EncodeToString(salt)
	nonce, err := hex.DecodeString(str)
	if err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New())

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plain, err := aesgcm.Open(nil, nonce, cipherText[:len(cipherText)-12], nil)

	dstFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	_, err = dstFile.Write(plain)
	if err != nil {
		panic(err.Error())
	}
}
