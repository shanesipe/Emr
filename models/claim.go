type Claim struct {
  ID string
  PatientID string
  TotalCost float64
  LineItems []ClaimLineItem
}

type ClaimLineItem struct {
  ID string
  Code string
  Cost float64
}
