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
