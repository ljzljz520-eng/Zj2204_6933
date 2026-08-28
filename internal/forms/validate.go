package forms

import (
	"encoding/json"
	"errors"
	"hospitalforms/internal/model"
)

func ValidatePayload(r model.Record) error {
	if len(r.Payload) == 0 {
		return errors.New("empty payload")
	}
	var m map[string]string
	if e := json.Unmarshal([]byte(r.Payload), &m); e != nil {
		return e
	}
	for _, d := range model.DefaultDefinitions() {
		if d.Type == r.FormType && len(d.Validate(m)) > 0 {
			return errors.New("payload violates form policy")
		}
	}
	return nil
}
func NormalizeStatus(status string) string {
	switch status {
	case "draft", "processed", "archived":
		return status
	default:
		return "draft"
	}
}
