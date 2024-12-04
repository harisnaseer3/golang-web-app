package handlers

import (
	"encoding/json"
	"net/http"
	"golang-web-app/models"
	"golang-web-app/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByUsername(creds.Username)
	if err != nil || !utils.CheckPasswordHash(creds.Password, user.PasswordHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
