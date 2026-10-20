package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manger struct {
	globalMiddlewares []Middleware
}

func NewManger() *Manger {
	return &Manger{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manger) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manger) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i]
	// 	n = middleware(n)
	// }

	// for _, middleware := range slices.Backward(middlewares) {
	// 	n = middleware(n)
	// }

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}

	return n
}
