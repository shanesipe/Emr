type DoctorService interface {
  GetDoctor(id string) (*Doctor, error)
  // other methods
}

type DefaultDoctorService struct {
  repo DoctorRepository 
}

// Implement service methods
