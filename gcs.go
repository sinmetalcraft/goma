package goma

import (
	"context"
	"fmt"
	"image"

	"cloud.google.com/go/storage"
	"github.com/disintegration/imaging"
)

type StorageService struct {
	gcs *storage.Client
}

func NewStorageService(ctx context.Context, gcs *storage.Client) *StorageService {
	return &StorageService{gcs: gcs}
}

func (s *StorageService) Read(ctx context.Context, bucket string, object string) (image.Image, string, error) {
	attrs, err := s.gcs.Bucket(bucket).Object(object).Attrs(ctx)
	if err != nil {
		return nil, "", err
	}

	r, err := s.gcs.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		if err := r.Close(); err != nil {
			fmt.Printf("failed GCS Reader Close(). err=%+v\n", err)
		}
	}()
	dst, err := imaging.Decode(r)
	if err != nil {
		return nil, "", err
	}

	return dst, attrs.ContentType, nil
}

func (s *StorageService) Write(ctx context.Context, img *image.NRGBA, format imaging.Format, bucket string, object string) (rerr error) {
	w := s.gcs.Bucket(bucket).Object(object).NewWriter(ctx)
	defer func() {
		if err := w.Close(); err != nil {
			if rerr == nil {
				rerr = err
			} else {
				fmt.Printf("failed GCS Writer Close(). err=%+v\n", err)
			}
		}
	}()
	if err := imaging.Encode(w, img, format); err != nil {
		return err
	}

	return nil
}
