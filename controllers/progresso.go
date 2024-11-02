package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"course-management-api/models"
	"course-management-api/database"
)

func CreateProgresso(w http.ResponseWriter, r *http.Request) {
	var novoProgresso models.Progresso
	if err := json.NewDecoder(r.Body).Decode(&novoProgresso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&novoProgresso).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoProgresso)
}

func GetProgresso(w http.ResponseWriter, r *http.Request) {
	var progresso []models.Progresso
	if err := database.DB.Find(&progresso).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(progresso)
}

func UpdateProgresso(w http.ResponseWriter, r *http.Request) {
	var updatedProgresso models.Progresso
	if err := json.NewDecoder(r.Body).Decode(&updatedProgresso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var progresso models.Progresso
	if err := database.DB.First(&progresso, updatedProgresso.UsuarioID, updatedProgresso.CursoID).Error; err != nil {
		http.Error(w, "Progresso não encontrado", http.StatusNotFound)
		return
	}

	if err := database.DB.Save(&updatedProgresso).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProgresso)
}

func DeleteProgresso(w http.ResponseWriter, r *http.Request) {
	usuarioIDStr := r.URL.Query().Get("usuario_id")
	cursoIDStr := r.URL.Query().Get("curso_id")

	usuarioID, err := strconv.Atoi(usuarioIDStr)
	if err != nil {
		http.Error(w, "ID de usuário inválido", http.StatusBadRequest)
		return
	}

	cursoID, err := strconv.Atoi(cursoIDStr)
	if err != nil {
		http.Error(w, "ID de curso inválido", http.StatusBadRequest)
		return
	}

	var progresso models.Progresso
	if err := database.DB.First(&progresso, usuarioID, cursoID).Error; err != nil {
		http.Error(w, "Progresso não encontrado", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&progresso).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
