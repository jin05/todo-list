package middleware

import (
	"context"
	"log"
	"net/http"
)

type Middlewares interface {
	List() []func(http.Handler) http.Handler
}

type middlewares struct {
	middlewares []func(http.Handler) http.Handler
}

type Handler func(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error)

func NewMiddlewares(
	auth AuthMiddleware,
	cors CORSMiddleware,
) Middlewares {
	m := &middlewares{}
	m.middlewares = append(m.middlewares, m.Middleware(auth.Handler))
	m.middlewares = append(m.middlewares, m.Middleware(cors.Handler))
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
				err = errorForContext(r.Context())
				if err != nil {
					log.Println(err)
					http.Error(w, "InternalServerError", http.StatusInternalServerError)
				}
			}
		})
	}
}

type errorContextKey struct{}

var errorKey = errorContextKey{}

func SetError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorKey, err)
}

func errorForContext(ctx context.Context) error {
	row, _ := ctx.Value(errorKey).(error)
	return row
}
