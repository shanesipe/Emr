func main() {

  // Load config
  config := LoadConfig()
  
  // Initialize database
  db := InitializeDB(config)

  // Initialize repositories
  patientRepo := NewPatientRepository(db)

  // Initialize services
  patientService := NewPatientService(patientRepo)

  // Initialize API routes
  router := routes.NewRouter()
  router.AddHandler("/patients", patientHandlers.GetPatients)

  // Start HTTP server
  http.ListenAndServe(":8080", router)
}
