package goma

import (
	"image"
	"io"

	"github.com/disintegration/imaging"
)

func Write(w io.Writer, img image.Image, format FormatType) error {
	var f imaging.Format
	switch format {
	case PNG:
		f = imaging.PNG
	case JPEG:
		f = imaging.JPEG
	default:
		f = imaging.PNG
	}

	return imaging.Encode(w, img, f)
}
