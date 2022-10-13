package middlewares

import (
	"net/http"
	// "task/models"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// func SetMiddlewareAuthentication() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		err := ValidateSign(r)
// 		if err != nil {
// 			models.ErrorResponse(w, 401, err)
// 			return
// 		}
// 		next(w, r)
// 	}
// }
