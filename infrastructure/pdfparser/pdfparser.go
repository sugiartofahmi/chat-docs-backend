package pdfparser

import (
	"bytes"
	"io"

	"github.com/ledongthuc/pdf"
)

func ExtractText(data []byte) (string, error) {
	reader, err := pdf.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return "", err
	}

	textReader, err := reader.GetPlainText()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, textReader); err != nil {
		return "", err
	}

	return buf.String(), nil
}
