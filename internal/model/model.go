package model

import "time"

type Record struct {
	ID, PatientID, FormType, Status, Payload string
	CreatedAt, UpdatedAt                     time.Time
	Version                                  int
}
type Profile struct {
	ID, Name, Department, Contact string
	Active                        bool
	UpdatedAt                     time.Time
}
type Event struct {
	ID, RecordID, Kind, Actor, Detail string
	At                                time.Time
}
type Audit struct {
	ID, RecordID, Action, Actor, Result string
	At                                  time.Time
}
type ArchiveEntry struct {
	RecordID, Reason, ArchivedBy string
	ArchivedAt                   time.Time
}
type Query struct {
	PatientID, FormType, Status string
	Limit                       int
}

func (r Record) Valid() bool      { return r.ID != "" && r.PatientID != "" && r.FormType != "" }
func (r Record) IsDraft() bool    { return r.Status == "draft" }
func (r Record) IsArchived() bool { return r.Status == "archived" }
func NewRecord(id, patient, form, payload string) Record {
	now := time.Now().UTC()
	return Record{ID: id, PatientID: patient, FormType: form, Payload: payload, Status: "draft", CreatedAt: now, UpdatedAt: now, Version: 1}
}
