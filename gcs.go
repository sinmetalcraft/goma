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

type WithStorageWriteOption func(ctx context.Context, oh *storage.ObjectHandle, w *storage.Writer) error

func WithMaxAge(age int) WithStorageWriteOption {
	return func(ctx context.Context, oh *storage.ObjectHandle, w *storage.Writer) error {
		w.CacheControl = fmt.Sprintf("public,max-age=%d", age)
		return nil
	}
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

func (s *StorageService) Write(ctx context.Context, img *image.NRGBA, format imaging.Format, bucket string, object string, ops ...WithStorageWriteOption) (rerr error) {
	oh := s.gcs.Bucket(bucket).Object(object)
	w := oh.NewWriter(ctx)
	defer func() {
		if err := w.Close(); err != nil {
			if rerr == nil {
				rerr = err
			} else {
				fmt.Printf("failed GCS Writer Close(). err=%+v\n", err)
			}
		}
	}()
	for _, op := range ops {
		if err := op(ctx, oh, w); err != nil {
			return err
		}
	}

	if err := imaging.Encode(w, img, format); err != nil {
		return err
	}

	return nil
}

func (s *StorageService) AddObjectACL(ctx context.Context, bucket string, object string, entity storage.ACLEntity, role storage.ACLRole) error {
	stg, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	roles, err := stg.Bucket(bucket).Object(object).ACL().List(ctx)
	if err != nil {
		return err
	}
	roles = append(roles, storage.ACLRule{Entity: entity, Role: role})
	_, err = stg.Bucket(bucket).Object(object).Update(ctx, storage.ObjectAttrsToUpdate{ACL: roles})
	if err != nil {
		return err
	}

	return nil
}
