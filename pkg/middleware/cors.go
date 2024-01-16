package middleware

import "net/http"

func enableCors(w http.ResponseWriter) http.ResponseWriter {
	// w.Header().Set("Access-Control-Allow-Origin", "https://apex-new-site.lavina.tech, https://apexpizza.uz")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	return w
}

// Cors enables cors policy
func Cors(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w = enableCors(w)
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, req)
		}
	})
}
