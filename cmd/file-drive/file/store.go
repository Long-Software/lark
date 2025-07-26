package file

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type StoreOpts struct {
	Root              string
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	defaultFolder := "file-network"
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if len(opts.Root) == 0 {
		opts.Root = defaultFolder
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) WriteStream(key string, r io.Reader) error {
	kPath := s.PathTransformFunc(key)
	if err := os.MkdirAll(s.Path(kPath), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(s.Path(kPath))
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

	f, err := os.Open(s.Path(kPath))
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
	return os.RemoveAll(s.Root + "/"+ kPath.Root())
}

func (s *Store) Has(key string) bool {
	kPath := s.PathTransformFunc(key)

	_, err := os.Stat(kPath.Root())
	return err == nil
}

func (s *Store) Path(kPath KeyPath) string {
	return fmt.Sprintf("%s/%s", s.Root, kPath.Path())
}
