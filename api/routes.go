func SetupRoutes(router *mux.Router, services *Services) {

  // Patient routes
  router.HandleFunc("/patients", patient.GetPatientsHandler(services.PatientService)).Methods("GET")
  router.HandleFunc("/patients/{id}", patient.GetPatientHandler(services.PatientService)).Methods("GET")
  
  // Doctor routes
  router.HandleFunc("/doctors", doctor.GetDoctorsHandler(services.DoctorService)).Methods("GET")
  
  // Appointment routes
  router.HandleFunc("/appointments", appointment.GetAppointmentsHandler(services.AppointmentService)).Methods("GET")
  
  // Lab routes
  router.HandleFunc("/lab/orders", lab.CreateLabOrderHandler(services.LabService)).Methods("POST")
  
  // Medication routes
  router.HandleFunc("/medications", medication.GetMedicationsHandler(services.MedicationService)).Methods("GET")

  // Billing routes
  router.HandleFunc("/claims", billing.SubmitClaimHandler(services.BillingService)).Methods("POST")

}
