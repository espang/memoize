package memoize

import (
	"container/list"
)

func M(f func(int) int) func(int) int {
	state := make(map[int]int)
	return func(v int) int {
		if r, ok := state[v]; ok {
			return r
		}

		r := f(v)
		state[v] = r
		return r
	}
}

func MFiFo(f func(int) int) func(int) int {
	state := make(map[int]int)
	queue := list.New()
	capacity := 10

	return func(v int) int {
		if r, ok := state[v]; ok {
			return r
		}

		if queue.Len() == capacity {
			evict := queue.Remove(queue.Front()).(int)
			delete(state, evict)
		}

		r := f(v)
		state[v] = r
		_ = queue.PushBack(v)
		return r
	}
}

func MLRU(f func(int) int) func(int) int {
	state := make(map[int]int)
	elements := make(map[int]*list.Element)
	lru := list.New()
	capacity := 10

	return func(v int) int {
		if elem, ok := elements[v]; ok {
			lru.MoveToBack(elem)
			return state[v]
		}

		if lru.Len() == capacity {
			evict := lru.Remove(lru.Front()).(int)
			delete(state, evict)
			delete(elements, evict)
		}

		r := f(v)
		state[v] = r
		elements[v] = lru.PushBack(v)
		return r
	}
}
