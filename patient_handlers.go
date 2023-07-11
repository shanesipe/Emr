package handlers

import (
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"

  "emr/models"
  "emr/services"
)

// PatientHandlers handles patient API requests
type PatientHandlers struct {
  service services.PatientService
}

// GetPatients returns a list of patients
func (h *PatientHandlers) GetPatients(w http.ResponseWriter, r *http.Request) {
  
  // Get filters from query params
  filters := r.URL.Query()

  patients, err := h.service.GetPatients(filters)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // Encode patients as JSON
  json.NewEncoder(w).Encode(patients)
}

// GetPatient returns a patient by ID
func (h *PatientHandlers) GetPatient(w http.ResponseWriter, r *http.Request) {

  // Get id from request
  id := mux.Vars(r)["id"]

  patient, err := h.service.GetPatient(id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }

  json.NewEncoder(w).Encode(patient)
}

// CreatePatient creates a new patient
func (h *PatientHandlers) CreatePatient(w http.ResponseWriter, r *http.Request) {

  // Decode request body into patient struct
  var patient models.Patient
  json.NewDecoder(r.Body).Decode(&patient)

  err := h.service.CreatePatient(&patient)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  json.NewEncoder(w).Encode(patient)
}

// UpdatePatient updates a patient
func (h *PatientHandlers) UpdatePatient(w http.ResponseWriter, r *http.Request) {

  // Get ID from request
  id := mux.Vars(r)["id"]

  // Decode request body
  var patient models.Patient
  json.NewDecoder(r.Body).Decode(&patient)

  patient.ID = id // Set ID from request params

  err := h.service.UpdatePatient(&patient)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  json.NewEncoder(w).Encode(patient)  
}
