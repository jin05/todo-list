package middleware

import "net/http"

type Middlewares interface {
	List() []func(http.Handler) http.Handler
}

type middlewares struct {
	middlewares []func(http.Handler) http.Handler
}

type Handler func(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error)

func NewMiddlewares(
	auth AuthMiddleware,
) Middlewares {
	m := &middlewares{}
	m.middlewares = append(m.middlewares, m.Middleware(auth.Handler))
	return m
}

func (m *middlewares) List() []func(http.Handler) http.Handler {
	return m.middlewares
}

func (m *middlewares) Middleware(handler Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w, r, err := handler(w, r)
			if err == nil {
				next.ServeHTTP(w, r)
			}
		})
	}
}
