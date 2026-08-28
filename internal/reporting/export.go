package reporting

import (
	"encoding/json"
	"hospitalforms/internal/model"
)

func JSON(records []model.Record) ([]byte, error) { return json.MarshalIndent(records, "", "  ") }
func GroupByPatient(records []model.Record) map[string][]model.Record {
	m := map[string][]model.Record{}
	for _, r := range records {
		m[r.PatientID] = append(m[r.PatientID], r)
	}
	return m
}
