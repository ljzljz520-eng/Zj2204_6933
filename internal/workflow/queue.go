package workflow

import (
	"context"
	"hospitalforms/internal/model"
	"sync"
)

type Queue struct {
	mu    sync.Mutex
	items []model.Record
}

func (q *Queue) Push(r model.Record) { q.mu.Lock(); defer q.mu.Unlock(); q.items = append(q.items, r) }
func (q *Queue) Pop(ctx context.Context) (model.Record, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if err := ctx.Err(); err != nil {
		return model.Record{}, false
	}
	if len(q.items) == 0 {
		return model.Record{}, false
	}
	r := q.items[0]
	q.items = q.items[1:]
	return r, true
}
func (q *Queue) Len() int { q.mu.Lock(); defer q.mu.Unlock(); return len(q.items) }
