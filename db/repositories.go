package db

import "emr/models"

// PatientRepository provides patient data access
type PatientRepository interface {
  GetPatient(id string) (*models.Patient, error)
  GetPatients(filters Filters) ([]*models.Patient, error)
  CreatePatient(p *models.Patient) error
  UpdatePatient(p *models.Patient) error
}

// DoctorRepository provides doctor data access
type DoctorRepository interface {
  // Similar methods for doctor model
}
