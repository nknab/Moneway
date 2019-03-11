package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(BalanceService) BalanceService

type loggingMiddleware struct {
	logger log.Logger
	next   BalanceService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BalanceService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next BalanceService) BalanceService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetBalance(ctx context.Context, accountID string) (string, error) {
	defer func() {
		l.logger.Log("method", "GetBalance", "accountID", accountID)
	}()
	return l.next.GetBalance(ctx, accountID)
}
func (l loggingMiddleware) UpdateBalance(ctx context.Context, accountID string, amount string) (string, error) {
	defer func() {
		l.logger.Log("method", "UpdateBalance", "accountID", accountID, "amount", amount)
	}()
	return l.next.UpdateBalance(ctx, accountID, amount)
}
