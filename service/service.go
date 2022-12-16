package service

import "github.com/dedeyuyandi/go-template/repository"

type gbTemplateService struct {
	redis repository.RedisReadWriter
}

func NewTemplateService(
	redis repository.RedisReadWriter,
) gbTemplateService {
	return gbTemplateService{
		redis: redis,
	}
}
