package helper

import "net/http"

func EnableCors(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(w).Header().Set("Content-Security-Policy", "default-src 'self'")
} //mengatur sumber daya
