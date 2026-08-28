package reporting

import (
	"context"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/model"
	"sort"
)

type Summary struct{ Total, Draft, Processed, Archived int }

func Summarize(ctx context.Context, s *forms.Service) (Summary, error) {
	rs, e := s.Query(ctx, model.Query{})
	if e != nil {
		return Summary{}, e
	}
	out := Summary{Total: len(rs)}
	for _, r := range rs {
		switch r.Status {
		case "draft":
			out.Draft++
		case "processed":
			out.Processed++
		case "archived":
			out.Archived++
		}
	}
	return out, nil
}
func SortByUpdated(rs []model.Record) {
	sort.Slice(rs, func(i, j int) bool { return rs[i].UpdatedAt.Before(rs[j].UpdatedAt) })
}
