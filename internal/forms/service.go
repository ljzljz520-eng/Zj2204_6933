package forms

import (
	"context"
	"errors"
	"fmt"
	"hospitalforms/internal/model"
	"hospitalforms/internal/storage"
	"time"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) Register(ctx context.Context, r model.Record) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !r.Valid() {
		return errors.New("invalid record")
	}
	return s.Store.SaveRecord(r)
}
func (s *Service) Process(ctx context.Context, id, actor string) (model.Event, error) {
	if err := ctx.Err(); err != nil {
		return model.Event{}, err
	}
	r, e := s.Store.FindRecord(id)
	if e != nil {
		return model.Event{}, e
	}
	if !r.IsDraft() {
		return model.Event{}, errors.New("record not draft")
	}
	if e = s.Store.UpdateStatus(id, "processed"); e != nil {
		return model.Event{}, e
	}
	v := model.Event{ID: fmt.Sprintf("evt-%d", time.Now().UnixNano()), RecordID: id, Kind: "processed", Actor: actor, At: time.Now().UTC()}
	return v, s.Store.SaveEvent(v)
}
func (s *Service) Archive(ctx context.Context, id, actor, reason string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if _, e := s.Store.FindRecord(id); e != nil {
		return e
	}
	if e := s.Store.UpdateStatus(id, "archived"); e != nil {
		return e
	}
	return s.Store.SaveArchive(model.ArchiveEntry{RecordID: id, Reason: reason, ArchivedBy: actor, ArchivedAt: time.Now().UTC()})
}
func (s *Service) Restore(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.Store.UpdateStatus(id, "draft")
}
func (s *Service) Query(ctx context.Context, q model.Query) ([]model.Record, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.Store.ListRecords(q)
}
