type LabOrder struct {
  ID string
  PatientID string
  TestCodes []string
  DateOrdered time.Time
}

type LabResult struct {
  ID string
  LabOrderID string
  Results map[string]interface{}
}
