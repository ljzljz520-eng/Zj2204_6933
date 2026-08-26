package model

type FieldRule struct {
	Name                string
	Required, Sensitive bool
	MaxLength           int
}
type FormDefinition struct {
	Type    string
	Version int
	Rules   []FieldRule
}

func DefaultDefinitions() []FormDefinition {
	return []FormDefinition{{Type: "triage", Version: 1, Rules: []FieldRule{{Name: "chief_complaint", Required: true, MaxLength: 500}, {Name: "allergies", Sensitive: true, MaxLength: 1000}}}, {Type: "discharge", Version: 1, Rules: []FieldRule{{Name: "diagnosis", Required: true, MaxLength: 1000}, {Name: "instructions", Required: true, MaxLength: 2000}}}}
}
func (d FormDefinition) Validate(payload map[string]string) []string {
	var out []string
	for _, r := range d.Rules {
		v, ok := payload[r.Name]
		if r.Required && !ok {
			out = append(out, r.Name+" required")
			continue
		}
		if ok && r.MaxLength > 0 && len(v) > r.MaxLength {
			out = append(out, r.Name+" too long")
		}
	}
	return out
}
