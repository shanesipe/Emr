type Medication struct {
  ID string
  Name string
  Strength string 
  Code string
}

type MedicationOrder struct {
  ID string
  PatientID string
  MedicationID string
  Dosage string
}
