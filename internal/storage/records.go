package storage

import (
	"go.etcd.io/bbolt"
	"hospitalforms/internal/model"
	"time"
)

func (s *Store) SaveRecord(r model.Record) error {
	if r.UpdatedAt.IsZero() {
		r.UpdatedAt = time.Now().UTC()
	}
	b, e := model.EncodeRecord(r)
	if e != nil {
		return e
	}
	return s.put("records", r.ID, b)
}
func (s *Store) FindRecord(id string) (*model.Record, error) {
	b, e := s.get("records", id)
	if e != nil {
		return nil, e
	}
	r, e := model.DecodeRecord(b)
	if e != nil {
		return nil, e
	}
	return &r, nil
}
func (s *Store) ListRecords(q model.Query) ([]model.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			r, er := model.DecodeRecord(v)
			if er != nil {
				return er
			}
			if q.PatientID != "" && r.PatientID != q.PatientID {
				return nil
			}
			if q.FormType != "" && r.FormType != q.FormType {
				return nil
			}
			if q.Status != "" && r.Status != q.Status {
				return nil
			}
			out = append(out, r)
			if q.Limit > 0 && len(out) >= q.Limit {
				return nil
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) UpdateStatus(id, status string) error {
	r, e := s.FindRecord(id)
	if e != nil {
		return e
	}
	r.Status = status
	r.UpdatedAt = time.Now().UTC()
	r.Version++
	return s.SaveRecord(*r)
}
