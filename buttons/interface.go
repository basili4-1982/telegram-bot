package buttons

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

//go:generate mockgen -source ./interface.go -destination=mocks/service.go -package=mocks
type Storage interface {
	AddRate(ctx context.Context, userID string, rate float64) error
}

type Rate interface {
	Get(symbol string) (float64, error)
}

type TeleContex interface {
	tele.Context
}
