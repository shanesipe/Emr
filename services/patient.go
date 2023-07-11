package models

import "time"

// Patient represents a patient record
type Patient struct {
  ID        string     `json:"id"`
  Name      string     `json:"name"` 
  DateOfBirth time.Time `json:"dateOfBirth"`
  Gender    string     `json:"gender"`
  Address   string     `json:"address"`
  Phone     string     `json:"phone"`
  Email     string     `json:"email"`
}

// TableName sets Patient's table name to be `patients`
func (Patient) TableName() string {
  return "patients" 
}

// PatientSummary is a summarized view of a patient
type PatientSummary struct {
  ID        string     `json:"id"`
  Name      string     `json:"name"`
  DateOfBirth time.Time `json:"dateOfBirth"`
  Gender    string     `json:"gender"` 
}

// TableName sets PatientSummary's table name to be `patient_summaries`
func (PatientSummary) TableName() string {
  return "patient_summaries"
}
