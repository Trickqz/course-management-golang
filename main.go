package main

import (
    "log"
    "net/http"
    "course-management-api/routes"
)

func main() {
    r := routes.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080", r))
}