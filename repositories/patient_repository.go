package repositories

import "emr/models"

// PatientRepository provides access to patient storage
type PatientRepository interface {
  GetPatient(id string) (*models.Patient, error)
  SavePatient(patient *models.Patient) error
}

// InMemoryPatientRepo implements PatientRepository using in-memory storage
type InMemoryPatientRepo struct {
  patients map[string]*models.Patient 
}

// GetPatient fetches a patient by ID
func (r *InMemoryPatientRepo) GetPatient(id string) (*models.Patient, error) {
  // implementation
} 

// SavePatient stores a new patient
func (r *InMemoryPatientRepo) SavePatient(patient *models.Patient) error {
  // implementation  
}
