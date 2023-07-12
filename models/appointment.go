package models

import (
    "time"

    "github.com/google/uuid"
)

// Appointment represents a doctor appointment booked by a patient
type Appointment struct {
    ID        string    `json:"id"`
    PatientID string    `json:"patient_id"`
    DoctorID  string    `json:"doctor_id"`
    Date      time.Time `json:"date"`
    Notes     string    `json:"notes"`
}

// NewAppointment creates a new appointment
func NewAppointment(patientID, doctorID string, date time.Time, notes string) *Appointment {
    return &Appointment{
        ID:        uuid.New().String(),
        PatientID: patientID,
        DoctorID:  doctorID,
        Date:      date,
        Notes:     notes,
    }
}

// IsValid validates the appointment
func (a *Appointment) IsValid() bool {
    return a.PatientID != "" && a.DoctorID != "" && !a.Date.IsZero()
}

// GetPatientID returns the patient ID
func (a *Appointment) GetPatientID() string {
    return a.PatientID
}

// GetDoctorID returns the doctor ID  
func (a *Appointment) GetDoctorID() string {
    return a.DoctorID
}

// GetDate returns the appointment date
func (a *Appointment) GetDate() time.Time {
    return a.Date
}

// GetNotes returns the appointment notes
func (a *Appointment) GetNotes() string {
    return a.Notes
}

// SetNotes sets the appointment notes
func (a *Appointment) SetNotes(notes string) {
    a.Notes = notes
}

