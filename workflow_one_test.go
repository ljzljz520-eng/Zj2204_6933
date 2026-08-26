package main

import (
	"context"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/model"
	"hospitalforms/internal/storage"
	"hospitalforms/internal/workflow"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	e := workflow.New(forms.New(s))
	if err := e.Intake(context.Background(), model.NewRecord("a", "p", "triage", "{}"), "nurse"); err != nil {
		t.Fatal(err)
	}
}
