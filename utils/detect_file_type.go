package utils

import (
	"io"

	"github.com/gabriel-vasile/mimetype"
)

func DetectFileType(file io.Reader) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	fileType := mimetype.Detect(buffer)
	return fileType.Extension(), nil
}
