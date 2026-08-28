package main

import (
	"context"
	"hospitalforms/internal/model"
	"hospitalforms/internal/workflow"
	"testing"
)

func TestQueueLifecycle(t *testing.T) {
	q := &workflow.Queue{}
	q.Push(model.NewRecord("r", "p", "f", ""))
	if q.Len() != 1 {
		t.Fatal(q.Len())
	}
	if _, ok := q.Pop(context.Background()); !ok || q.Len() != 0 {
		t.Fatal("queue")
	}
}
