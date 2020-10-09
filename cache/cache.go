package cache


type Cache[Item any] interface {
	Put(key string, item Item) error
	Get(key string) (Item, bool)
	BatchGet(keys ...string) (map[string]Item)
}

type LoadingFunc[Item any] func(string) (Item, bool)

func New[Item any]() Cache[Item] {
	return nil
}

func NewLoading[Item any](fn LoadingFunc[Item]) Cache[Item] {
	
}
