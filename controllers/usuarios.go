package controllers

import (
	"encoding/json"
	"net/http"
	"sync"
	"strconv"
	"course-management-api/models"
	"course-management-api/database"
	"golang.org/x/crypto/bcrypt"
)

var (
	usuarios  = []models.Usuario{}
	usuariosMu sync.Mutex
)

func Registrar(w http.ResponseWriter, r *http.Request) {
	var novoUsuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&novoUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(novoUsuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao criar senha", http.StatusInternalServerError)
		return
	}
	novoUsuario.Senha = string(hashedPassword)

	if err := database.DB.Create(&novoUsuario).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoUsuario)
}

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario
	if err := database.DB.Find(&usuarios).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	var updatedUsuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&updatedUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	if err := database.DB.First(&usuario, updatedUsuario.ID).Error; err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	if err := database.DB.Save(&updatedUsuario).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUsuario)
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	if err := database.DB.First(&usuario, id).Error; err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&usuario).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
