package models

import "time"

// Claim represents an insurance claim
type Claim struct {
    ID string
    PatientID string
    AppointmentID string
    Date time.Time
    Amount float64
    Status string
}

// NewClaim creates a new claim
func NewClaim(patientID, appointmentID string, amount float64) *Claim {
    return &Claim{
        ID: uuid.New().String(),
        PatientID: patientID,
        AppointmentID: appointmentID,
        Date: time.Now(),
        Amount: amount,
        Status: "Pending",
    }
}

// GetID returns the claim ID
func (c *Claim) GetID() string {
    return c.ID
}

// GetPatientID returns the patient ID
func (c *Claim) GetPatientID() string {
    return c.PatientID
}

// GetAppointmentID returns the appointment ID
func (c *Claim) GetAppointmentID() string {
    return c.AppointmentID
}

// GetDate returns the date of the claim
func (c *Claim) GetDate() time.Time {
    return c.Date
}

// GetAmount returns the claim amount
func (c *Claim) GetAmount() float64 {
    return c.Amount
}

// GetStatus returns the status of the claim
func (c *Claim) GetStatus() string {
    return c.Status
}

// SetStatus sets the status of the claim
func (c *Claim) SetStatus(status string) {
    c.Status = status
}
