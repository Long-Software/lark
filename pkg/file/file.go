package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Create(path string) error {
	dirPath := filepath.Dir(path)
	err := MkdirAll(dirPath)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
func Append(path string, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}
func Write(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
func Copy(from string, to string) error {
	srcFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	err = Create(to)
	destFile, err := os.Open(to)
	if err != nil {
		return err
	}
	defer destFile.Close()
	// Copy the file content
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// Optionally, copy file permissions
	srcInfo, err := os.Stat(from)
	if err == nil {
		err = os.Chmod(to, srcInfo.Mode())
	}

	return err
}

func Delete(file_path string) error {
	return os.Remove(file_path)
}

func GetExecDir() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	execDir := filepath.Dir(execPath)
	return execDir, nil
}

func ListFilesWithExtension(dirPath string, ext string) ([]string, error) {
	var files []string

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ext {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

func GetFileInfo(filePath string) (os.FileInfo, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	// Ensure it's a file, not a directory
	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory, not a file", filePath)
	}
	return info, err
}
