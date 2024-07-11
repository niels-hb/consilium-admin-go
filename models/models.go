package models

import (
	"encoding/json"
	"time"
)

type FileExport struct {
	Transactions []TransactionExport `json:"transactions"`
	Schedules    []ScheduleExport    `json:"schedules"`
}

type TransactionExport struct {
	UID         string     `json:"-"`
	AmountCents int        `json:"amount_cents"`
	Category    string     `json:"category"`
	CreatedOn   *time.Time `json:"created_on"`
	Name        string     `json:"name"`
	Note        string     `json:"note,omitempty"`
}

func (t *TransactionExport) FromJSON(data map[string]interface{}) {
	marshalled, _ := json.Marshal(data)

	var export TransactionExport
	json.Unmarshal(marshalled, &export)

	*t = export
}

func (t *TransactionExport) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid":          t.UID,
		"amount_cents": t.AmountCents,
		"category":     t.Category,
		"created_on":   t.CreatedOn,
		"name":         t.Name,
		"note":         t.Note,
	}
}

type ScheduleExport struct {
	UID             string     `json:"-"`
	AmountCents     int        `json:"amount_cents"`
	CanceledOn      *time.Time `json:"canceled_on,omitempty"`
	Category        string     `json:"category"`
	CreatedOn       *time.Time `json:"created_on"`
	FrequencyMonths int        `json:"frequency_months"`
	Name            string     `json:"name"`
	Note            string     `json:"note,omitempty"`
	StartedOn       *time.Time `json:"started_on"`
	ScheduleType    string     `json:"type"`
}

func (s *ScheduleExport) FromJSON(data map[string]interface{}) {
	marshalled, _ := json.Marshal(data)

	var export ScheduleExport
	json.Unmarshal(marshalled, &export)

	*s = export
}

func (s *ScheduleExport) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid":              s.UID,
		"amount_cents":     s.AmountCents,
		"canceled_on":      s.CanceledOn,
		"category":         s.Category,
		"created_on":       s.CreatedOn,
		"frequency_months": s.FrequencyMonths,
		"name":             s.Name,
		"note":             s.Note,
		"started_on":       s.StartedOn,
		"type":             s.ScheduleType,
	}
}
