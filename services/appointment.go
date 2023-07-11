type AppointmentService interface {
  GetAppointment(id string) (*Appointment, error)
  CreateAppointment(a *Appointment) error
  // other methods
}

// Implement service
