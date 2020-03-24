package goma

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/disintegration/imaging"
)

func Write(w io.Writer, img image.Image, ft FormatType) error {
	return imaging.Encode(w, img, ft.ImagingFormat())
}

func Open(path string, ft FormatType) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	switch ft {
	case JPEG:
		img, err := jpeg.Decode(file)
		if err != nil {
			return nil, err
		}

		return img, nil
	case PNG:
		img, err := png.Decode(file)
		if err != nil {
			return nil, err
		}

		return img, nil
	default:
		return nil, fmt.Errorf("%v is unsupported", ft)
	}
}
