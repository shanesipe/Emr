type MedicationService interface {
  GetMedication(id string) (*Medication, error)
  CreateMedicationOrder(order *MedicationOrder) error
}

// Implement service
