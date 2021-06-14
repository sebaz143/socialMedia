package middleware

import (
	"net/http"

	"github.com/sebaz143/socialMedia/bd"
)

/*ChequeoBD es un pasa manos que me permite revisar el estado de la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexion Perdida cno la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
