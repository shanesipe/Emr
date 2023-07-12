package services

import (
  "errors"

  "github.com/jinzhu/gorm"

  "emr/models"
  "emr/repositories"
)

// PatientService provides operations on patients
type PatientService interface {
  GetPatient(id string) (*models.Patient, error)
  GetPatients(filters Filters) ([]*models.Patient, error)
  CreatePatient(patient *models.Patient) error
  UpdatePatient(patient *models.Patient) error
}

// DefaultPatientService implements PatientService
type DefaultPatientService struct {
  repo repositories.PatientRepository
}

// GetPatient returns a patient by ID
func (s *DefaultPatientService) GetPatient(id string) (*models.Patient, error) {
  return s.repo.FindByID(id)
}

// GetPatients returns a filtered list of patients
func (s *DefaultPatientService) GetPatients(filters Filters) ([]*models.Patient, error) {
  return s.repo.FindAll(filters) 
}

// CreatePatient creates a new patient
func (s *DefaultPatientService) CreatePatient(patient *models.Patient) error {
  err := s.repo.Create(patient)
  if err != nil {
    return err
  }
  return nil
}

// UpdatePatient updates an existing patient
func (s *DefaultPatientService) UpdatePatient(patient *models.Patient) error {
  err := s.repo.Update(patient)
  if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return ErrPatientNotFound
    }
    return err
  }
  return nil
}
