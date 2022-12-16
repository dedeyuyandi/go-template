package repository

type RedisReadWriter interface {
	io.Closer

	HSetOrderDispatchCache(key string, field string, value model.OrderDispatch) error
	// HGetOrderDispatchCache(key, field string) (resp model.OrderDispatch)
	// HGetAllOrderDispatchCache(key string) (resp []model.OrderDispatch)
}
