package services

import (
    "context"
    "errors"

    "emr/models"
    "emr/repositories"
)

// AppointmentService provides appointment management functions
type AppointmentService interface {
    GetAppointment(ctx context.Context, id string) (*models.Appointment, error)
    CreateAppointment(ctx context.Context, appointment *models.Appointment) error
    UpdateAppointment(ctx context.Context, appointment *models.Appointment) error
    DeleteAppointment(ctx context.Context, id string) error
}

type appointmentService struct {
    appointmentRepo repositories.AppointmentRepository
}

// GetAppointment retrieves an appointment by ID
func (s *appointmentService) GetAppointment(ctx context.Context, id string) (*models.Appointment, error) {
    appointment, err := s.appointmentRepo.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return appointment, nil
}

// CreateAppointment saves a new appointment
func (s *appointmentService) CreateAppointment(ctx context.Context, appointment *models.Appointment) error {
    err := s.appointmentRepo.Save(ctx, appointment)
    if err != nil {
        return err
    }
    return nil
}

// UpdateAppointment updates an existing appointment
func (s *appointmentService) UpdateAppointment(ctx context.Context, appointment *models.Appointment) error {
    // Validate appointment exists
    if appointment.ID == "" {
        return errors.New("invalid appointment")
    }
    
    return s.appointmentRepo.Update(ctx, appointment)
}

// DeleteAppointment removes an appointment
func (s *appointmentService) DeleteAppointment(ctx context.Context, id string) error {
    // Validate ID
    if id == "" {
        return errors.New("invalid ID")
    }
    
    return s.appointmentRepo.Delete(ctx, id)
}

// NewAppointmentService creates an AppointmentService
func NewAppointmentService(appointmentRepo repositories.AppointmentRepository) AppointmentService {
    return &appointmentService{appointmentRepo: appointmentRepo}
}
