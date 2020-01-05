package goma

import (
	"context"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/disintegration/imaging"
	"github.com/google/go-cmp/cmp"
)

func TestResize(t *testing.T) {
	s := newStorageService(t)

	ctx := context.Background()

	cases := []struct {
		name                   string
		srcObject              string
		dstObject              string
		wantGomaType           GomaType
		dstWidth               int
		dstHeight              int
		dstContentType         imaging.Format
		withStorageWriteOption []WithStorageWriteOption
	}{
		{"png", "shingo_nouhau.png", "shingo_nouhau_600.png", GomaType{PNG, "image/png"}, 600, 0, imaging.PNG, []WithStorageWriteOption{WithMaxAge(600)}},
		{"png2", "sinmetal-merpay.png", "sinmetal-merpay_600.png", GomaType{PNG, "image/png"}, 600, 0, imaging.PNG, []WithStorageWriteOption{WithMaxAge(600)}},
		{"jpeg", "shingo.jpg", "shingo_600.jpg", GomaType{JPEG, "image/jpeg"}, 600, 0, imaging.JPEG, []WithStorageWriteOption{WithMaxAge(600)}},
		{"no-change-size", "sinmetal.jpg", "sinmetal_640.jpg", GomaType{JPEG, "image/jpeg"}, 640, 0, imaging.JPEG, []WithStorageWriteOption{WithMaxAge(600)}},
		{"mini-to-mini", "sinmetal.jpg", "sinmetal_320.jpg", GomaType{JPEG, "image/jpeg"}, 320, 0, imaging.JPEG, []WithStorageWriteOption{WithMaxAge(600)}},
		{"no-cache", "sinmetal.jpg", "sinmetal_no_cache.jpg", GomaType{JPEG, "image/jpeg"}, 320, 0, imaging.JPEG, []WithStorageWriteOption{}},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			src, gt, err := s.Read(ctx, "sinmetal", tt.srcObject)
			if err != nil {
				t.Fatal(err)
			}

			if e, g := tt.wantGomaType, gt; cmp.Equal(tt.wantGomaType, gt) {
				t.Errorf("GomaType Type want %v got %v", e, g)
			}

			dst := Resize(src, tt.dstWidth, tt.dstHeight)
			if err := s.Write(ctx, dst, gt.FormatType, "sinmetal", tt.dstObject, tt.withStorageWriteOption...); err != nil {
				t.Fatal(err)
			}
			if err := s.AddObjectACL(ctx, "sinmetal", tt.dstObject, storage.AllUsers, storage.RoleReader); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestResizeToFitLongSide(t *testing.T) {
	s := newStorageService(t)

	ctx := context.Background()

	cases := []struct {
		name           string
		srcObject      string
		dstObject      string
		wantGomaType   GomaType
		dstSize        int
		dstContentType imaging.Format
	}{
		{"png", "shingo_nouhau.png", "shingo_nouhau_300.png", GomaType{PNG, "image/png"}, 300, imaging.PNG},
		{"jpeg", "shingo.jpg", "shingo_300.jpg", GomaType{JPEG, "image/jpeg"}, 300, imaging.JPEG},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			src, gt, err := s.Read(ctx, "sinmetal", tt.srcObject)
			if err != nil {
				t.Fatal(err)
			}
			if e, g := tt.wantGomaType, gt; cmp.Equal(e, g) {
				t.Errorf("Content Type want %v got %v", e, g)
			}

			dst := ResizeToFitLongSide(src, tt.dstSize)
			if err := s.Write(ctx, dst, gt.FormatType, "sinmetal", tt.dstObject); err != nil {
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
