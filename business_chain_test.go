package main

import (
	"context"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/storage"
	"path/filepath"
	"testing"
)

func TestBusinessChain06(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	f := forms.New(s)
	defer func() {
		if recover() != nil {
			t.Fatalf("missing form must return a user-facing error")
		}
	}()
	_, err := f.RecoverMissing(context.Background(), "missing")
	if err == nil {
		t.Fatal("missing form should be reported")
	}
}
