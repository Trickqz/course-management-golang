package controllers

import (
    "encoding/json"
    "net/http"
    "sync"
    "strconv"
    "course-management-api/models"
)

var (
    cursos  = []models.Curso{}
    cursosMu sync.Mutex
)

func CreateCurso(w http.ResponseWriter, r *http.Request) {
    var novoCurso models.Curso
    if err := json.NewDecoder(r.Body).Decode(&novoCurso); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    cursosMu.Lock()
    novoCurso.ID = len(cursos) + 1
    cursos = append(cursos, novoCurso)
    cursosMu.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(novoCurso)
}

func GetCursos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cursos)
}

func UpdateCurso(w http.ResponseWriter, r *http.Request) {
    var updatedCurso models.Curso
    if err := json.NewDecoder(r.Body).Decode(&updatedCurso); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    cursosMu.Lock()
    for i, curso := range cursos {
        if curso.ID == updatedCurso.ID {
            cursos[i] = updatedCurso
            cursosMu.Unlock()
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedCurso)
            return
        }
    }
    cursosMu.Unlock()
    http.Error(w, "Curso não encontrado", http.StatusNotFound)
}

func DeleteCurso(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    cursosMu.Lock()
    defer cursosMu.Unlock()

    for i, curso := range cursos {
        if curso.ID == id {
            cursos = append(cursos[:i], cursos[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "Curso não encontrado", http.StatusNotFound)
}