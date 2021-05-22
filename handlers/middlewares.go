package handlers

import (
	"go.uber.org/zap"
	"net/http"
)

type Middlewares struct {
	logger *zap.SugaredLogger
}

func NewMiddlewares(logger *zap.SugaredLogger) *Middlewares {
	return &Middlewares{logger: logger}
}

func (mw *Middlewares) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		mw.logger.Infow("request", "uri", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
