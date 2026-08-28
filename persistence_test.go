package main

import (
	"hospitalforms/internal/model"
	"hospitalforms/internal/storage"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveRecord(model.NewRecord("r1", "p1", "triage", "{}")); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, e := s.FindRecord("r1")
	if e != nil || r.ID != "r1" {
		t.Fatalf("reopen failed: %v", e)
	}
}
