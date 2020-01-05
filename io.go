package goma

import (
	"image"
	"io"

	"github.com/disintegration/imaging"
)

func Write(w io.Writer, img image.Image, ft FormatType) error {
	return imaging.Encode(w, img, ft.ImagingFormat())
}
