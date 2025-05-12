package sqlite

import (
	"binanceTemporal/entities"
	"context"
)

type Sqlite interface {
	AddKlineEvent(ctx context.Context, klineEvent *entities.KlineEvent) error
}
