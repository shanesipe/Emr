package models

import "time"

// Patient represents a medical patient
type Patient struct {
  ID string
  Name string
  DateOfBirth time.Time
  Gender string
  Address string
  Phone string
  Email string
  MedicalHistory []string
}

// NewPatient creates a new patient
func NewPatient(name string, dob time.Time, gender string) *Patient {
  return &Patient{
    ID: uuid.New().String(),
    Name: name,
    DateOfBirth: dob, 
    Gender: gender,
  }
}

// GetID returns the patient ID
func (p *Patient) GetID() string {
  return p.ID
}

// GetName returns the patient name
func (p *Patient) GetName() string {
  return p.Name
}

// GetDateOfBirth returns the DOB
func (p *Patient) GetDateOfBirth() time.Time {
  return p.DateOfBirth
}

// GetGender returns the gender
func (p *Patient) GetGender() string {
  return p.Gender
}

// SetAddress updates the address
func (p *Patient) SetAddress(address string) {
  p.Address = address
}

// AddMedicalHistory adds a history entry 
func (p *Patient) AddMedicalHistory(entry string) {
  p.MedicalHistory = append(p.MedicalHistory, entry)
}
