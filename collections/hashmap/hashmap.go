package hashmap

type HashMap[K, V any] struct {
	cmp func(T, T) int
}


func Make[K, V any](cmp func(K, K) int) HashMap[K, V] {
	return HashMap[K, V]{cmp: cmp}
}
