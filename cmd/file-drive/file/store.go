package file

import (
	"bytes"
	"io"
	"os"
)

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) WriteStream(key string, r io.Reader) error {
	kPath := s.PathTransformFunc(key)
	if err := os.MkdirAll(kPath.Folder, os.ModePerm); err != nil {
		return err
	}

	filePath := kPath.Path()

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ReadStream(key string) (io.Reader, error) {
	kPath := s.PathTransformFunc(key)

	f, err := os.Open(kPath.Path())
	if err != nil {
		return nil, err
	}

	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	return buf, err
}

func (s *Store) Delete(key string) error {
	kPath := s.PathTransformFunc(key)
	return os.RemoveAll(kPath.Path())
}

func (s *Store) Has(key string) bool {
	kPath := s.PathTransformFunc(key)

	_, err := os.Stat(kPath.Path())
	return err == nil
}
