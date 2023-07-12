package models

import "time"

// Medication represents a prescription medication
type Medication struct {
  ID string
  Name string
  Description string
  Dosage string
}

// NewMedication creates a new Medication
func NewMedication(name, desc, dosage string) *Medication {
  return &Medication{
    ID: uuid.New().String(),
    Name: name,
    Description: desc,
    Dosage: dosage,
  }
}

// GetID returns the medication ID
func (m *Medication) GetID() string {
  return m.ID
}

// GetName returns the medication name
func (m *Medication) GetName() string {
  return m.Name
} 

// GetDescription returns the description
func (m *Medication) GetDescription() string {
  return m.Description
}

// GetDosage returns the dosage instructions
func (m *Medication) GetDosage() string {
  return m.Dosage
}

// SetDosage updates the dosage instructions
func (m *Medication) SetDosage(dosage string) {
  m.Dosage = dosage
}

// Prescription represents a patient's prescription
type Prescription struct {
  ID string
  PatientID string
  MedicationID string
  DatePrescribed time.Time
  Dosage string
} 

// NewPrescription creates a new Prescription
func NewPrescription(patientID, medicationID string) *Prescription {
  // implementation
}
