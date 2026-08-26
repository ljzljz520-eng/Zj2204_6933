package storage

import (
	"fmt"
	"go.etcd.io/bbolt"
	"hospitalforms/internal/model"
)

func (s *Store) SaveProfile(p model.Profile) error {
	b, e := model.EncodeProfile(p)
	if e != nil {
		return e
	}
	return s.put("profiles", p.ID, b)
}
func (s *Store) SaveEvent(v model.Event) error {
	b, e := model.EncodeEvent(v)
	if e != nil {
		return e
	}
	return s.put("events", v.ID, b)
}
func (s *Store) SaveAudit(v model.Audit) error {
	b, e := model.EncodeAudit(v)
	if e != nil {
		return e
	}
	return s.put("audits", v.ID, b)
}
func (s *Store) SaveArchive(v model.ArchiveEntry) error {
	return s.put("archives", v.RecordID, []byte(fmt.Sprintf("%s|%s|%s", v.RecordID, v.Reason, v.ArchivedBy)))
}
func (s *Store) Count(bucket string) (int, error) {
	n := 0
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error { n = tx.Bucket([]byte(bucket)).Stats().KeyN; return nil })
	return n, e
}
