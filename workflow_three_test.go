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

func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	e := workflow.New(forms.New(s))
	r := model.NewRecord("c", "p", "triage", "{}")
	_ = e.Forms.Register(context.Background(), r)
	if err := e.Recover(context.Background(), "c"); err != nil {
		t.Fatal(err)
	}
}
