package main

import (
	"hospitalforms/internal/model"
	"testing"
)

func TestModelValidation(t *testing.T) {
	if model.NewRecord("", "p", "f", "").Valid() {
		t.Fatal("invalid record accepted")
	}
	if !model.NewRecord("r", "p", "f", "").IsDraft() {
		t.Fatal("draft status")
	}
}
