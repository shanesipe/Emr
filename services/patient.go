type PatientService interface {
  GetPatient(id string) (*Patient, error) 
  CreatePatient(p *Patient) error
  // other methods  
}

type DefaultPatientService struct {
  repo PatientRepository
}

func (s *DefaultPatientService) GetPatient(id string) (*Patient, error) {
  return s.repo.GetPatient(id)
} 
