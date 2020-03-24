package goma

import (
	"context"
	"image"
	"image/draw"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/tenntenn/nigari"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type TextService struct {
	BaseFont  font.Face
	EmojiFont font.Face // カラー絵文字には対応してないので、今のところ使ってない https://github.com/tenntenn/nigari/issues/1
}

func NewTextService(ctx context.Context, baseFont font.Face, emojiFont font.Face) *TextService {
	return &TextService{
		BaseFont:  baseFont,
		EmojiFont: emojiFont,
	}
}

func (s *TextService) Draw(dst draw.Image, spacing float64, width fixed.Int26_6, fg image.Image, text string, x int, y int) {
	d := &nigari.Drawer{
		Base:    s.BaseFont,
		Emoji:   s.EmojiFont,
		Spacing: spacing,
		Width:   width,
	}
	d.Draw(text, x, y, dst, fg)
}

func LoadFont(path string, size float64) (font.Face, error) {
	fontBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(f, &truetype.Options{
		Size:    size,
		DPI:     128,
		Hinting: font.HintingNone,
	}), nil
}
