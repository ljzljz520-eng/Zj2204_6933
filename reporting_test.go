package main

import (
	"context"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/model"
	"hospitalforms/internal/reporting"
	"hospitalforms/internal/storage"
	"path/filepath"
	"testing"
)

func TestReportingSummary(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	f := forms.New(s)
	_ = f.Register(context.Background(), model.NewRecord("r", "p", "f", "{}"))
	x, e := reporting.Summarize(context.Background(), f)
	if e != nil || x.Total != 1 {
		t.Fatal(x, e)
	}
}
