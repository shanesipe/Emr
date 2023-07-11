type LabService interface {
  CreateLabOrder(order *LabOrder) error
  GetLabResult(id string) (*LabResult, error)
}

// Implement service
