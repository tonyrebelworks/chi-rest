package middleware

import (
	"chi-rest/lib"
	"net/http"

	"github.com/globalsign/mgo"
)

// AppMiddlewareInterface ...
type AppMiddlewareInterface interface {
	RequestLoggerMiddleware(next http.Handler) http.Handler
	NotfoundMiddleware(next http.Handler) http.Handler
}

// AppMiddleware ...
type appMiddleware struct {
	DB  *mgo.Session
	Cfg lib.Config
}

// NewAppMiddleware ...
func NewAppMiddleware(db *mgo.Session, config lib.Config) AppMiddlewareInterface {
	return &appMiddleware{db, config}
}
