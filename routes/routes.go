package routes

import (
    "github.com/gorilla/mux"
    "course-management-api/controllers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/cursos", controllers.GetCursos).Methods("GET")
    r.HandleFunc("/cursos", controllers.CreateCurso).Methods("POST")
    r.HandleFunc("/cursos/{id}", controllers.UpdateCurso).Methods("PUT")
    r.HandleFunc("/cursos", controllers.DeleteCurso).Methods("DELETE")
    r.HandleFunc("/usuarios", controllers.GetUsuarios).Methods("GET")
    r.HandleFunc("/usuarios", controllers.CreateUsuario).Methods("POST")
    r.HandleFunc("/usuarios/{id}", controllers.UpdateUsuario).Methods("PUT")
    r.HandleFunc("/usuarios", controllers.DeleteUsuario).Methods("DELETE")
    return r
}