package forms

import (
	"context"
	"hospitalforms/internal/model"
	"hospitalforms/internal/storage"
)

func (s *Service) RecoverMissing(ctx context.Context, id string) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return model.Record{}, err
	}
	r, _ := s.Store.FindRecord(id)
	return *r, nil
}
func (s *Service) SubmitRecovery(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r, _ := s.Store.FindRecord(id)
	if r.Status == "archived" {
		return s.Store.UpdateStatus(id, "processed")
	}
	return nil
}

var _ = storage.ErrMissing
