package controllers

import (
	"encoding/json"
	"net/http"
	"sync"
	"strconv"
	"course-management-api/models"
)

var (
	usuarios  = []models.Usuario{}
	usuariosMu sync.Mutex
)

func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var novoUsuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&novoUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usuariosMu.Lock()
	novoUsuario.ID = len(usuarios) + 1
	usuarios = append(usuarios, novoUsuario)
	usuariosMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoUsuario)
}

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	var updatedUsuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&updatedUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usuariosMu.Lock()
	defer usuariosMu.Unlock()

	for i, usuario := range usuarios {
		if usuario.ID == updatedUsuario.ID {
			usuarios[i] = updatedUsuario
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUsuario)
			return
		}
	}
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	usuariosMu.Lock()
	defer usuariosMu.Unlock()

	for i, usuario := range usuarios {
		if usuario.ID == id {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}
