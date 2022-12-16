package repository

import (
	"io"

	"github.com/dedeyuyandi/go-template/model"
)

type RedisReadWriter interface {
	io.Closer

	HSetOrderDispatchCache(key string, field string, value model.OrderDispatch) error
	// HGetOrderDispatchCache(key, field string) (resp model.OrderDispatch)
	// HGetAllOrderDispatchCache(key string) (resp []model.OrderDispatch)
}
