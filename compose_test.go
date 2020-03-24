package goma

import (
	"context"
	"image"
	"testing"
)

func TestComposeCMYK(t *testing.T) {
	s := newStorageService(t)

	tis := openTestImages(t)

	var reqs []*ComposeImageRequest
	reqs = append(reqs, &ComposeImageRequest{
		img: tis.Background,
		x:   0,
		y:   0,
	})
	reqs = append(reqs, &ComposeImageRequest{
		img: tis.Red,
		x:   100,
		y:   100,
	})

	opts := []WithStorageWriteOption{WithMaxAge(600)}
	ctx := context.Background()
	img := ComposeCMYK(640, 480, reqs)
	if err := s.Write(ctx, img, PNG, "sinmetal", "compose-cmyk.png", opts...); err != nil {
		t.Fatal(err)
	}
}

func TestComposeRGBA(t *testing.T) {
	s := newStorageService(t)

	tis := openTestImages(t)

	var reqs []*ComposeImageRequest
	reqs = append(reqs, &ComposeImageRequest{
		img: tis.Background,
		x:   0,
		y:   0,
	})
	reqs = append(reqs, &ComposeImageRequest{
		img: tis.Red,
		x:   100,
		y:   100,
	})

	opts := []WithStorageWriteOption{WithMaxAge(600)}
	ctx := context.Background()
	img := ComposeRGBA(640, 480, reqs)
	if err := s.Write(ctx, img, PNG, "sinmetal", "compose-rgba.png", opts...); err != nil {
		t.Fatal(err)
	}
}

type TestImages struct {
	Background image.Image
	Red        image.Image
}

func openTestImages(t *testing.T) *TestImages {
	background, err := Open("./assets/background.png")
	if err != nil {
		t.Fatalf("failed background image open. err=%+v", err)
	}

	red, err := Open("./assets/red.png")
	if err != nil {
		t.Fatalf("failed red image open. err=%+v", err)
	}

	return &TestImages{
		Background: background,
		Red:        red,
	}
}
