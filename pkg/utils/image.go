package utils

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

func DecodeImage(file *os.File) (image.Image, string, error) {
	ext := strings.ToLower(filepath.Ext(file.Name()))
	switch ext {
	case ".jpg", ".jpeg":
		img, err := jpeg.Decode(file)
		return img, "JPG", err
	case ".png":
		img, err := png.Decode(file)
		return img, "PNG", err
	case ".gif":
		img, err := gif.Decode(file)
		return img, "GIF", err
	case ".bmp":
		img, err := bmp.Decode(file)
		return img, "BMP", err
	case ".tiff":
		img, err := tiff.Decode(file)
		return img, "TIFF", err
	case ".webp":
		img, err := webp.Decode(file)
		return img, "WEBP", err
	default:
		return nil, "", fmt.Errorf("unsupported image format: %s", ext)
	}
}

func IsImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp":
		return true
	}
	return false
}

func Dimension(r image.Rectangle) (width, height float64) {
	return float64(r.Dx()), float64(r.Dy())
}