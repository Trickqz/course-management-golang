package main

import (
	"log"
	"net/http"
	"course-management-api/database"
	"course-management-api/models"
	"course-management-api/routes"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(&models.Usuario{}, &models.Curso{}, &models.Progresso{})
	if err != nil {
		log.Fatalf("Erro ao migrar os modelos: %v", err)
	}

	r := routes.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}