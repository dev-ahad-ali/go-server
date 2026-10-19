package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manger struct {
	globalMiddlewares []Middleware
}

func NewManger() *Manger {
	return &Manger{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manger) With(middlewares ...Middleware) {

}
