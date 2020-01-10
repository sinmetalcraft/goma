package goma

import (
	"context"
	"image"
	"image/png"
	"os"
	"testing"

	"golang.org/x/image/math/fixed"
)

func TestTextService_DrawCMYK(t *testing.T) {
	ctx := context.Background()

	base, err := LoadFont("./assets/AozoraMinchoRegular.ttf", 10)
	if err != nil {
		t.Fatal(err)
	}
	emoji, err := LoadFont("./assets/NotoEmoji-Regular.ttf", 10)
	if err != nil {
		t.Fatal(err)
	}

	ts := NewTextService(ctx, base, emoji)

	dst := image.NewCMYK(image.Rect(0, 0, 300, 300))
	ts.DrawCMYK(dst, 1.5, fixed.I(270), image.Black, "Hello 絵文字じゃない文字列たち。突然の絵文字が襲いかかる。🍣🍺。", 10, 25)

	file, err := os.Create("./test/test.png")
	if err != nil {
		t.Fatal(err)
	}
	if err := png.Encode(file, dst); err != nil {
		t.Fatal(err)
	}
}
