package goma

import (
	"context"
	"testing"
)

func TestComposeCMYK(t *testing.T) {
	s := newStorageService(t)

	background, err := Open("./assets/background.png", PNG)
	if err != nil {
		t.Fatalf("failed background image open. err=%+v", err)
	}

	red, err := Open("./assets/red.png", PNG)
	if err != nil {
		t.Fatalf("failed red image open. err=%+v", err)
	}

	var reqs []*ComposeImageRequest
	reqs = append(reqs, &ComposeImageRequest{
		img: background,
		x:   0,
		y:   0,
	})
	reqs = append(reqs, &ComposeImageRequest{
		img: red,
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

	background, err := Open("./assets/background.png", PNG)
	if err != nil {
		t.Fatalf("failed background image open. err=%+v", err)
	}

	red, err := Open("./assets/red.png", PNG)
	if err != nil {
		t.Fatalf("failed red image open. err=%+v", err)
	}

	var reqs []*ComposeImageRequest
	reqs = append(reqs, &ComposeImageRequest{
		img: background,
		x:   0,
		y:   0,
	})
	reqs = append(reqs, &ComposeImageRequest{
		img: red,
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
