package file

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

type PathTransformFunc func(string) KeyPath

func CASPathTransformFunc(key string) KeyPath {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	slideLen := len(hashStr) / blockSize

	path := make([]string, slideLen)

	for i := 0; i < slideLen; i++ {
		from, to := i*blockSize, (i+1)*blockSize
		path[i] = hashStr[from:to]
	}

	return KeyPath{
		Folder:   strings.Join(path, "/"),
		Filename: hashStr,
	}
}

var DefaultPathTransformFunc = func(key string) KeyPath {
	return KeyPath{
		Folder: key,
	}
}