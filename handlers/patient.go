func GetPatientsHandler(service PatientService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    patients := service.GetPatients()
    json.NewEncoder(w).Encode(patients)
  }
}

func GetPatientHandler(service PatientService) http.HandlerFunc {

  return func(w http.ResponseWriter, r *http.Request) {
    
    id := mux.Vars(r)["id"]
    
    patient, err := service.GetPatient(id)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError) 
      return
    }

    json.NewEncoder(w).Encode(patient)
  }
}
