package routes

import (
    "github.com/gorilla/mux"
    "course-management-api/controllers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/cursos", controllers.GetCursos).Methods("GET")
    r.HandleFunc("/cursos", controllers.CreateCurso).Methods("POST")
    return r
}