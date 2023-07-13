package models

import "time"

type Appointment struct {
  ID string
  PatientID string
  DoctorID string 
  Date time.Time
  Notes string
}

func NewAppointment(patientID, doctorID string, date time.Time, notes string) *Appointment {
  return &Appointment{
    ID:       // generate ID
    PatientID: patientID,
    DoctorID: doctorID,
    Date: date,
    Notes: notes, 
  }
}

// Add other methods like GetID(), GetDate() etc
