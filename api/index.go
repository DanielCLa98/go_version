package handler

import (
    "net/http"
    "encoding/json"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // Configurar headers para JSON
    w.Header().Set("Content-Type", "application/json")
    
    // Crear respuesta
    response := map[string]string{
        "message": "¡Hola, Mundo desde Go!",
    }
    
    // Enviar JSON como respuesta
    json.NewEncoder(w).Encode(response)
}