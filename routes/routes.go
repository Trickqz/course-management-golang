package routes

import (
	"github.com/gorilla/mux"
	"course-management-api/controllers"
	"course-management-api/middleware"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/registrar", controllers.Registrar).Methods("POST")

	r.HandleFunc("/usuarios", middleware.VerificarToken(controllers.GetUsuarios)).Methods("GET")
	r.HandleFunc("/usuarios/{id}", middleware.VerificarToken(controllers.UpdateUsuario)).Methods("PUT")
	r.HandleFunc("/usuarios", middleware.VerificarToken(controllers.DeleteUsuario)).Methods("DELETE")

	r.HandleFunc("/cursos", middleware.VerificarToken(controllers.GetCursos)).Methods("GET")
	r.HandleFunc("/cursos", middleware.VerificarToken(controllers.CreateCurso)).Methods("POST")
	r.HandleFunc("/cursos/{id}", middleware.VerificarToken(controllers.UpdateCurso)).Methods("PUT")
	r.HandleFunc("/cursos", middleware.VerificarToken(controllers.DeleteCurso)).Methods("DELETE")

	r.HandleFunc("/progresso", middleware.VerificarToken(controllers.GetProgresso)).Methods("GET")
	r.HandleFunc("/progresso", middleware.VerificarToken(controllers.CreateProgresso)).Methods("POST")
	r.HandleFunc("/progresso", middleware.VerificarToken(controllers.UpdateProgresso)).Methods("PUT")
	r.HandleFunc("/progresso", middleware.VerificarToken(controllers.DeleteProgresso)).Methods("DELETE")

	return r
}