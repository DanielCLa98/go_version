package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola, Mundo desde Go en Vercel!")
}