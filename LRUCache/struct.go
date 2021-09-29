package LRUCache

type LRUCache struct {
    m map[int]int
	cap int
	stack []int
}


func Constructor(capacity int) LRUCache {
    return LRUCache {
		m: make(map[int]int),
		cap: capacity,
		stack: []int{},
	}
}


func (cache *LRUCache) Get(key int) int {
    v, ok := cache.m[key]
	if !ok {
		return -1
	}
	return v
}


func (cache *LRUCache) Put(key int, value int)  {
	_, ok := cache.m[key]
	if ok {
		cache.m[key] = value
		return
	}
    if len(cache.stack) >= cache.cap {
		n := len(cache.stack) - 1 
		delete(cache.m, cache.stack[n])
		cache.stack = cache.stack[:n]
	}
	cache.m[key] = value;
	cache.stack = append(cache.stack, key)
}