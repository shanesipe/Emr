package services

import (
  "context"

  "emr/models"
  "emr/repositories"
)

// BillingService generates invoices for appointments
type BillingService interface {
  GenerateInvoice(ctx context.Context, appointmentID string) (*models.Invoice, error)
}

type billingService struct {
  appointmentRepo repositories.AppointmentRepository
}

// GenerateInvoice creates an invoice for an appointment
func (b *billingService) GenerateInvoice(ctx context.Context, appointmentID string) (*models.Invoice, error) {
  // Get appointment
  appointment, err := b.appointmentRepo.GetAppointment(ctx, appointmentID)
  if err != nil {
    return nil, err
  }

  // Create invoice
  invoice := models.NewInvoice(appointment.PatientID, appointment.DoctorID, appointment.Date)

  // Calculate amount
  // Apply discounts

  // Save invoice
  err = b.invoiceRepo.Save(ctx, invoice)
  if err != nil {
    return nil, err 
  }

  return invoice, nil
}

// NewBillingService creates a BillingService 
func NewBillingService(appointmentRepo repositories.AppointmentRepository, 
                       invoiceRepo repositories.InvoiceRepository) BillingService {
  return &billingService{
    appointmentRepo: appointmentRepo,
    invoiceRepo: invoiceRepo,
  }
}

