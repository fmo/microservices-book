package ports

import (
	"context"
	"github.com/fmo/microservices-book/order/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(ctx context.Context, order *domain.Order) error
}
