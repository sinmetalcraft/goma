package goma

import (
	"context"
	"net/http/httptest"
	"testing"
)

func TestWrite(t *testing.T) {
	s := newStorageService(t)

	ctx := context.Background()

	src, _, err := s.Read(ctx, "sinmetal", "shingo_nouhau.png")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	if err := Write(w, src, PNG); err != nil {
		t.Fatal(err)
	}
}
