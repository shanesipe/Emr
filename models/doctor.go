package models

import "time"

// Doctor represents a medical doctor
type Doctor struct {
  ID string
  Name string
  Specialty string
  Experience int // years of experience
}

// NewDoctor creates a new Doctor
func NewDoctor(name, specialty string, experience int) *Doctor {
  return &Doctor{
    ID: uuid.New().String(),
    Name: name, 
    Specialty: specialty,
    Experience: experience,
  }
}

// GetID returns the doctor ID
func (d *Doctor) GetID() string {
  return d.ID
}

// GetName returns the doctor's name
func (d *Doctor) GetName() string {
  return d.Name
}

// GetSpecialty returns doctor's specialty
func (d *Doctor) GetSpecialty() string {
  return d.Specialty
}

// GetExperience returns years of experience
func (d *Doctor) GetExperience() int {
  return d.Experience
}

// SetSpecialty updates doctor's specialty
func (d *Doctor) SetSpecialty(specialty string) {
  d.Specialty = specialty
}

// IncrementExperience increments doctor's experience by 1 year
func (d *Doctor) IncrementExperience() {
  d.Experience++
}
