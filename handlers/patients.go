package handlers

import (
	"encoding/json"
	"net/http"
	"golang-web-app/models"
	"golang-web-app/middleware"
)

func GetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := models.GetAllPatients()
	if err != nil {
		http.Error(w, "Error fetching patients", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(patients)
}

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	if !middleware.IsReceptionist(r) {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := models.CreatePatient(&patient); err != nil {
		http.Error(w, "Error creating patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}
