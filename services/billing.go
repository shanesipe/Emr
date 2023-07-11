type BillingService interface {
  SubmitClaim(claim *Claim) error
  GetClaimStatus(id string) (*ClaimStatus, error) 
}

// Implement service
