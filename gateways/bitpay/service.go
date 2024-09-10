package bitpay

import (
	"fmt"

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
)

type BitPayIR struct {
	Logger *zap.Logger
	API    string
}

type BitPayRequest struct {
	OrderID     string
	Name        string
	Email       string
	Description string
	CallbackURL string
	Amount      int
}

func NewService(token string) (*BitPayIR, error) {
	err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Error closing response body: %w", zap.Error(err))
		}
	}()

	return &BitPayIR{
		API:    token,
		Logger: logger,
	}, nil
}
