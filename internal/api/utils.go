package api

import "net/http"

func corsHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Configurar cabeçalhos CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permitir qualquer origem
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			return
		}

		fn(w, r)
	}
}
