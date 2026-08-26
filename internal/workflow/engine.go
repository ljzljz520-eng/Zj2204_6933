package workflow

import (
	"context"
	"errors"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/model"
)

type Engine struct{ Forms *forms.Service }

func New(f *forms.Service) *Engine { return &Engine{Forms: f} }
func (e *Engine) Intake(ctx context.Context, r model.Record, actor string) error {
	if err := e.Forms.Register(ctx, r); err != nil {
		return err
	}
	_, err := e.Forms.Process(ctx, r.ID, actor)
	return err
}
func (e *Engine) Complete(ctx context.Context, id, actor string) error {
	if _, err := e.Forms.Process(ctx, id, actor); err != nil {
		return err
	}
	return e.Forms.Archive(ctx, id, actor, "completed")
}
func (e *Engine) Recover(ctx context.Context, id string) error {
	if err := e.Forms.SubmitRecovery(ctx, id); err != nil {
		return err
	}
	return nil
}
func ValidateChain(r model.Record) ([]string, error) {
	if !r.Valid() {
		return nil, errors.New("invalid")
	}
	return []string{"receive", "validate", "save", "display"}, nil
}
