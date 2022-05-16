package main

// hash map
// under the hood, a hash map uses a dynamic array of linked lists to efficiently store key-value pairs
// when inserting a new key-value pair, a hash function first map the key to an integer value, it will be used for index in the underlying dynamic array
// the map checks if the key already exists. if it does, it overwrites the old value with the new one
// below is an example of a hash map look like under the hood:
// [
//	0: [key1, value1] -> null,
//	1: [key2, value2] -> [key3, value3] -> null,
//	3: [key4, value4] -> null,
//	4: [key5, value5] -> [key6, value6] -> [key7, value7] -> null,
// ]
// in the hash map above, the keys 2, 3 collided by all being hashed to 1, and the keys 4, 5, 6, 7 collided by all being hashed to 4
// - insert -> O(1) on average; O(n) on worst case
// - lookup -> O(1) on average; O(n) on worst case
// - delete -> O(1) on average; O(n) on worst case
// - iterate -> O(n) on average; O(n) on worst case
// - size -> O(1)

// just use native map from golang
type HashMap[T any] map[string]T

func NewHashMap[T any]() HashMap[T] {
	return make(map[string]T)
}

func (h *HashMap[T]) Insert(key string, value T) {
	(*h)[key] = value
}

func (h *HashMap[T]) Get(key string) T {
	return (*h)[key]
}

func (h *HashMap[T]) Delete(key string) {
	delete(*h, key)
}

func (h *HashMap[T]) Size() int {
	return len(*h)
}

func (h *HashMap[T]) Iterate(f func(key string, value T)) {
	for k, v := range *h {
		f(k, v)
	}
}
