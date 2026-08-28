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

func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	e := workflow.New(forms.New(s))
	r := model.NewRecord("b", "p", "triage", "{}")
	if err := e.Forms.Register(context.Background(), r); err != nil {
		t.Fatal(err)
	}
	if err := e.Complete(context.Background(), "b", "doctor"); err != nil {
		t.Fatal(err)
	}
}
