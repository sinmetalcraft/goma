package goma

import (
	"context"
	"image"
	"image/png"
	"os"
	"testing"
)

func TestTextService_DrawCMYK(t *testing.T) {
	ctx := context.Background()

	base, err := LoadFont("./assets/NotoSerifCJKjp-hinted/NotoSerifCJKjp-Regular.ttf", 10)
	if err != nil {
		t.Fatal(err)
	}
	emoji, err := LoadFont("./assets/NotoEmoji-Regular.ttf", 10)
	if err != nil {
		t.Fatal(err)
	}

	ts := NewTextService(ctx, base, emoji)

	dst := image.NewCMYK(image.Rect(0, 0, 200, 200))
	ts.DrawCMYK(dst, 1.5, 100, image.Black, "HelloWorld", 0, 0)

	file, err := os.Create("./test/test.png")
	if err != nil {
		t.Fatal(err)
	}
	if err := png.Encode(file, dst); err != nil {
		t.Fatal(err)
	}
}
