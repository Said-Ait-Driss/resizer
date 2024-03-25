package service

import (
	"bytes"
	"image/jpeg"

	"github.com/nfnt/resize"
)

func ResizeImage(data []byte, width, height int) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, resizedImg, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
