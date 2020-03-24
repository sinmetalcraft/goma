package goma

import (
	"context"
	"image"
	"image/draw"
	"os"
	"testing"

	"github.com/disintegration/imaging"
	"golang.org/x/image/math/fixed"
)

func TestTextService_Draw(t *testing.T) {
	ctx := context.Background()

	base, err := LoadFont("./assets/AozoraMinchoRegular.ttf", 10)
	if err != nil {
		t.Fatal(err)
	}
	emoji, err := LoadFont("./assets/NotoEmoji-Regular.ttf", 10)
	if err != nil {
		t.Fatal(err)
	}
	tis := openTestImages(t)

	ts := NewTextService(ctx, base, emoji)

	dst := tis.Background.(draw.Image)
	ts.Draw(dst, 1.5, fixed.I(270), image.Black, "Hello 絵文字じゃない文字列たち。突然の絵文字が襲いかかる。🍣🍺。", 10, 25)
	file, err := os.Create("./test/test.png")
	if err != nil {
		t.Fatal(err)
	}
	if err := imaging.Encode(file, dst, imaging.PNG); err != nil {
		t.Fatal(err)
	}
}
