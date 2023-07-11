package db

import (
  "emr/models"

  "github.com/jinzhu/gorm"
)

// PostgresPatientRepository implements PatientRepository
type PostgresPatientRepository struct {
  DB *gorm.DB
}

// GetPatient queries patient by ID
func (r *PostgresPatientRepository) GetPatient(id string) (*models.Patient, error) {
  var patient models.Patient
  if err := r.DB.First(&patient, "id = ?", id).Error; err != nil {
    return nil, err
  }
  return &patient, nil
}

// Implements other repository methods...

// PostgresDoctorRepository implements DoctorRepository

// Implements doctor repository methods...
