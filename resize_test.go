package goma

import (
	"context"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/disintegration/imaging"
)

func TestResize(t *testing.T) {
	s := newStorageService(t)

	ctx := context.Background()

	cases := []struct {
		name            string
		srcObject       string
		dstObject       string
		wantContentType string
		dstWidth        int
		dstHeight       int
		dstContentType  imaging.Format
	}{
		{"png", "shingo_nouhau.png", "shingo_nouhau_600.png", "image/png", 600, 0, imaging.PNG},
		{"png2", "sinmetal-merpay.png", "sinmetal-merpay_600.png", "image/png", 600, 0, imaging.PNG},
		{"jpeg", "shingo.jpg", "shingo_600.jpg", "image/jpeg", 600, 0, imaging.JPEG},
		{"no-change-size", "sinmetal.jpg", "sinmetal_640.jpg", "image/jpeg", 640, 0, imaging.JPEG},
		{"mini-to-mini", "sinmetal.jpg", "sinmetal_320.jpg", "image/jpeg", 320, 0, imaging.JPEG},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			src, ct, err := s.Read(ctx, "sinmetal", tt.srcObject)
			if err != nil {
				t.Fatal(err)
			}
			if e, g := tt.wantContentType, ct; e != g {
				t.Errorf("Content Type want %v got %v", e, g)
			}

			dst := Resize(src, tt.dstWidth, tt.dstHeight)
			if err := s.Write(ctx, dst, tt.dstContentType, "sinmetal", tt.dstObject); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func newStorageService(t *testing.T) *StorageService {
	ctx := context.Background()

	gcs, err := storage.NewClient(ctx)
	if err != nil {
		t.Fatal(err)
	}

	return NewStorageService(ctx, gcs)
}
