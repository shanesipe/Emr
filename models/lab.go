package models

import "time"

// Lab represents a medical lab facility
type Lab struct {
  ID string
  Name string
  Address string
}

// NewLab creates a new Lab 
func NewLab(name, address string) *Lab {
  return &Lab{
    ID: uuid.New().String(),
    Name: name,
    Address: address,
  }
}

// GetID returns the lab ID
func (l *Lab) GetID() string {
  return l.ID
} 

// GetName returns the lab name
func (l *Lab) GetName() string {
  return l.Name
}

// GetAddress returns the lab address
func (l *Lab) GetAddress() string {
  return l.Address
}

// SetAddress updates the lab address
func (l *Lab) SetAddress(address string) {
  l.Address = address
}

// LabTest represents a medical test conducted by a lab
type LabTest struct {
  ID string
  LabID string
  PatientID string
  TestType string
  Date time.Time
  Result string
}

// NewLabTest creates a new lab test
func NewLabTest(labID, patientID, testType string) *LabTest {
  return &LabTest{
    ID: uuid.New().String(),
    LabID: labID,
    PatientID: patientID,
    TestType: testType,
    Date: time.Now(),
  }
}
