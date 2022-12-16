package service

import (
	"context"

	"github.com/dedeyuyandi/go-template/model"
)

type TemplateService interface {
	HealthCheck(ctx context.Context) (model.OrderDispatch, error)
}
