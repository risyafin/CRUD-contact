package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("Lecang"), nil
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		_, ok := token.Claims.(*MyClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token invalid"))
			return
		}
		next(w, r)
	}

}

// import "net/http"

// type customMux struct {
// 	http.ServeMux
// 	middlewares []func(http.Handler) http.Handler
// }

// func (c *customMux) RegistrasiMiddleware(next func(next http.Handler) http.Handler) {
// 	c.middlewares = append(c.middlewares, next)
// }

// func (c *customMux) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// 	var current http.Handler = &c.ServeMux

// 	for _, next := range c.middlewares {
// 		current = next(current)
// 	}
// 	current.ServeHTTP(w, r)
// }
