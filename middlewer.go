package main

import "net/http"

type customMux struct {
	http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (c *customMux) RegistrasiMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *customMux) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}
