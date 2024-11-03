package controllers

import (
	"encoding/json"
	"net/http"
	"course-management-api/models"
	"course-management-api/database"
	"course-management-api/auth"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var usuarioDB models.Usuario
	if err := database.DB.Where("email = ?", usuario.Email).First(&usuarioDB).Error; err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuarioDB.Senha), []byte(usuario.Senha)); err != nil {
		http.Error(w, "Senha incorreta", http.StatusUnauthorized)
		return
	}

	token, err := auth.GerarToken(usuarioDB.ID)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
