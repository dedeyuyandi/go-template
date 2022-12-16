package service

import (
	"github.com/dedeyuyandi/go-template"
)

type gbTemplateService struct {
	redis repository.RedisConfiguration
}

func NewTemplateService(
	redis repository.Redis,
) gbTemplateService {
	return gbOrderStreamingService{
		redis: redis,
	}
}
