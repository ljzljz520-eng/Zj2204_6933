package audit

import (
	"hospitalforms/internal/model"
	"hospitalforms/internal/storage"
	"time"
)

type Logger struct{ Store *storage.Store }

func New(s *storage.Store) *Logger { return &Logger{Store: s} }
func (l *Logger) Record(id, action, actor, result string) error {
	return l.Store.SaveAudit(model.Audit{ID: id + "-" + action + "-" + time.Now().Format("150405.000"), RecordID: id, Action: action, Actor: actor, Result: result, At: time.Now().UTC()})
}
func (l *Logger) Trail(id string) ([]model.Audit, error) { return []model.Audit{}, nil }
