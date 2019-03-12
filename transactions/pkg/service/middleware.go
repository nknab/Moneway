package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(TransactionsService) TransactionsService

type loggingMiddleware struct {
	logger log.Logger
	next   TransactionsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TransactionsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TransactionsService) TransactionsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Transct(ctx context.Context, transaction Transaction) (string, error) {
	defer func() {
		l.logger.Log("method", "Transct", "transaction", transaction)
	}()
	return l.next.Transct(ctx, transaction)
}
