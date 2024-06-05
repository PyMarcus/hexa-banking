package app

import (
	"encoding/json"
	"net/http"

	"github.com/PyMarcus/banking_auth/dto"
	"github.com/PyMarcus/banking_auth/service"
)

type AuthHandler struct {
	s service.IAuthService
}

func (ah AuthHandler) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		thereIsAToken := r.Header.Get("token")
		if len(thereIsAToken) == 0 {
			writeResponse(w, http.StatusUnauthorized, "Not authorized")
			return
		} else {
			if ah.s.ValidateToken(&thereIsAToken){
				next.ServeHTTP(w, r)
			}else{
				writeResponse(w, http.StatusUnauthorized, "Not authorized")
				return 
			}
		}
	}
}

func (ah AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Unexpected data in json body")
	} else {
		response, err := ah.s.Login(request)
		if err != nil {
			writeResponse(w, http.StatusForbidden, err.Message)
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}

func (ah AuthHandler) Pv(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, 200, "Ok")
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
