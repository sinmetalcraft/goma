package goma

import (
	"image"

	"github.com/disintegration/imaging"
)

// Resize is 画像をリサイズを行う
// ex. width:500, height:0 アスペクト比を保ったまま、width:500でリサイズする
func Resize(src image.Image, width, height int) *image.NRGBA {
	return imaging.Resize(src, width, height, imaging.Lanczos)
}

// ResizeToFitLongSide is 長辺が指定したサイズになるようにアスペクト比を維持してリサイズする
func ResizeToFitLongSide(src image.Image, size int) *image.NRGBA {
	if src.Bounds().Size().X > src.Bounds().Size().Y {
		return imaging.Resize(src, size, 0, imaging.Lanczos)
	} else {
		return imaging.Resize(src, 0, size, imaging.Lanczos)
	}
}
