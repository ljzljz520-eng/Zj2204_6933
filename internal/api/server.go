package api

import (
	"context"
	"encoding/json"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/model"
	"net/http"
)

type Server struct{ Forms *forms.Service }

func New(f *forms.Service) *Server { return &Server{Forms: f} }
func (s *Server) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	var v model.Record
	if json.NewDecoder(r.Body).Decode(&v) != nil {
		http.Error(w, "bad json", 400)
		return
	}
	if e := s.Forms.Register(r.Context(), v); e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (s *Server) List(w http.ResponseWriter, r *http.Request) {
	rs, e := s.Forms.Query(context.Background(), model.Query{PatientID: r.URL.Query().Get("patient")})
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	_ = json.NewEncoder(w).Encode(rs)
}
