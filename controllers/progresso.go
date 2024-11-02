package controllers

import (
	"encoding/json"
	"net/http"
	"sync"
	"strconv"
	"course-management-api/models"
)

var (
	progresso  = []models.Progresso{}
	progressoMu sync.Mutex
)

func CreateProgresso(w http.ResponseWriter, r *http.Request) {
	var novoProgresso models.Progresso
	if err := json.NewDecoder(r.Body).Decode(&novoProgresso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	progressoMu.Lock()
	progresso = append(progresso, novoProgresso)
	progressoMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoProgresso)
}

func GetProgresso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(progresso)
}

func UpdateProgresso(w http.ResponseWriter, r *http.Request) {
	var updatedProgresso models.Progresso
	if err := json.NewDecoder(r.Body).Decode(&updatedProgresso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	progressoMu.Lock()
	defer progressoMu.Unlock()

	for i, p := range progresso {
		if p.UsuarioID == updatedProgresso.UsuarioID && p.CursoID == updatedProgresso.CursoID {
			progresso[i] = updatedProgresso
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProgresso)
			return
		}
	}
	http.Error(w, "Progresso não encontrado", http.StatusNotFound)
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

	progressoMu.Lock()
	defer progressoMu.Unlock()

	for i, p := range progresso {
		if p.UsuarioID == usuarioID && p.CursoID == cursoID {
			progresso = append(progresso[:i], progresso[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Progresso não encontrado", http.StatusNotFound)
}
