package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"course-management-api/models"
	"course-management-api/database"
)

func CreateCurso(w http.ResponseWriter, r *http.Request) {
	var novoCurso models.Curso
	if err := json.NewDecoder(r.Body).Decode(&novoCurso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&novoCurso).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoCurso)
}

func GetCursos(w http.ResponseWriter, r *http.Request) {
	var cursos []models.Curso
	if err := database.DB.Find(&cursos).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cursos)
}

func UpdateCurso(w http.ResponseWriter, r *http.Request) {
	var updatedCurso models.Curso
	if err := json.NewDecoder(r.Body).Decode(&updatedCurso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var curso models.Curso
	if err := database.DB.First(&curso, updatedCurso.ID).Error; err != nil {
		http.Error(w, "Curso não encontrado", http.StatusNotFound)
		return
	}

	if err := database.DB.Save(&updatedCurso).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCurso)
}

func DeleteCurso(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var curso models.Curso
	if err := database.DB.First(&curso, id).Error; err != nil {
		http.Error(w, "Curso não encontrado", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&curso).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}