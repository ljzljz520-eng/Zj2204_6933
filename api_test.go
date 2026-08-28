package main

import (
	"hospitalforms/internal/api"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/storage"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	rr := httptest.NewRecorder()
	api.New(forms.New(s)).Health(rr, httptest.NewRequest("GET", "/health", nil))
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}
