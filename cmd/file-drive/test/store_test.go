package test

import (
	"bytes"
	"io"
	"testing"

	"github.com/Long-Software/lark/cmd/file-drive/file"
	"github.com/Long-Software/lark/cmd/file-drive/utils"
	"github.com/Long-Software/lark/pkg/log"
)

func TestStore(t *testing.T) {
	s := file.NewStore(file.StoreOpts{
		PathTransformFunc: file.CASPathTransformFunc,
	})
	s.PathTransformFunc("")
	key := "pictures"
	data := "some text"

	if err := s.WriteStream(key, bytes.NewReader([]byte(data))); err != nil {
		t.Error(err)
	}

	r, err := s.ReadStream(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}

}

func TestPathTransformFunc(t *testing.T) {
	key := "some photos of me"
	pathName := file.CASPathTransformFunc(key)
	utils.Log.NewLog(log.DEBUG, pathName.Path())
}

func TestStoreDelete(t *testing.T) {
	s := file.NewStore(file.StoreOpts{
		PathTransformFunc: file.CASPathTransformFunc,
	})
	key := "deleted folder"
	data := []byte("some old text files")

	if err := s.WriteStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

