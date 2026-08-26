package storage

import (
	"go.etcd.io/bbolt"
	"hospitalforms/internal/model"
	"time"
)

type Batch struct {
	Records []model.Record
	Events  []model.Event
}

func (s *Store) SaveBatch(b Batch) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error {
		rb := tx.Bucket([]byte("records"))
		for _, r := range b.Records {
			x, e := model.EncodeRecord(r)
			if e != nil {
				return e
			}
			if e = rb.Put([]byte(r.ID), x); e != nil {
				return e
			}
		}
		eb := tx.Bucket([]byte("events"))
		for _, v := range b.Events {
			x, e := model.EncodeEvent(v)
			if e != nil {
				return e
			}
			if e = eb.Put([]byte(v.ID), x); e != nil {
				return e
			}
		}
		return nil
	})
}
func (s *Store) Touch(id string) error {
	r, e := s.FindRecord(id)
	if e != nil {
		return e
	}
	r.UpdatedAt = time.Now().UTC()
	return s.SaveRecord(*r)
}
