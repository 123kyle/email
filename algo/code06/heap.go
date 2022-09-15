package code06

import "errors"

type GreaterHeap[T comparable] struct {
	Limit      int
	Size       int
	Map        map[T]int
	Comparator func(a, b T) int
}

func NewGreaterHeap[T comparable](comparator func(a, b T) int, limit int) *GreaterHeap[T] {
	return &GreaterHeap[T]{
		Data:       make([]T, limit),
		Limit:      limit,
		Size:       0,
		Comparator: comparator,
		Map:        make(map[T]int),
	}
}

func (h *GreaterHeap[T]) Contains(obj T) bool {
	_, ok := h.Map[obj]
	return ok
}

func (h *GreaterHeap[T]) IsEmpty() bool {
	return h.Size == 0
}

func (h *GreaterHeap[T]) IsFull() bool {
	return h.Limit == h.Size
}

func (h *GreaterHeap[T]) Push(obj T) error {
	if h.IsFull() {
		return errors.New("heap is full")
	}
	index := h.Size
	h.Data[index] = obj
	h.Map[obj] = index
	h.Size++
	return h.heapInsert(index)
}

func (h *GreaterHeap[T]) Pop() (T, error) {
	var target T
	if h.IsEmpty() {
		return target, errors.New("heap is empty")
	}
	target = h.Data[0]

	//h.Data[h.Size-1], h.Data[0] = h.Data[0], h.Data[h.Size-1]
	h.swap(h.Size-1, 0)
	h.Size--
	delete(h.Map, target)
	var err error
	if h.Size != 0 {
		err = h.heapify(0)
	}
	return target, err
}

func (h *GreaterHeap[T]) Remove(obj T) {
	var replace T
	replace = h.Data[h.Size-1]
	index, ok := h.Map[obj]
	if !ok {
		return
	}
	h.swap(h.Size-1, index)
	//h.Data = append(h.Data[:index], h.Data[index+1:]...)
	h.Size--
	if obj != replace {
		h.Data[index] = replace
		h.Map[replace] = index
		h.Resign(replace)
	}
}

func (h *GreaterHeap[T]) Peek() (T, error) {
	var target T
	if h.IsEmpty() {
		return target, errors.New("heap is empty")
	}
	return h.Data[0], nil
}

func (h *GreaterHeap[T]) heapInsert(index int) error {
	for h.Comparator(h.Data[index], h.Data[(index-1)/2]) > 0 { // h.Data[index] > h.Data[(index-1)/2]
		//h.Data[index], h.Data[(index-1)/2] = h.Data[(index-1)/2], h.Data[index]
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
	return nil
}

func (h *GreaterHeap[T]) heapify(index int) error {
	if index >= h.Size {
		return errors.New("index out of range")
	}
	left := 2*index + 1
	right := 2*index + 2
	var bigger int
	for {
		if right < h.Size {
			if h.Comparator(h.Data[left], h.Data[right]) > 0 { // h.Data[left] > h.Data[right]
				bigger = left
			} else {
				bigger = right
			}
		} else if left < h.Size {
			bigger = left
		} else {
			break
		}
		if h.Comparator(h.Data[bigger], h.Data[index]) > 0 { // h.Data[bigger] > h.Data[index]
			//h.Data[bigger], h.Data[index] = h.Data[index], h.Data[bigger]
			h.swap(bigger, index)
			index = bigger
		} else {
			break
		}
		left = 2*index + 1
		right = 2*index + 2
	}
	return nil
}

func (h *GreaterHeap[T]) Resign(obj T) {
	index, ok := h.Map[obj]
	if ok {
		h.heapInsert(index)
		h.heapify(index)
	}
}

func (h *GreaterHeap[T]) swap(i, j int) {
	o1 := h.Data[i]
	o2 := h.Data[j]
	h.Data[i] = o2
	h.Data[j] = o1
	h.Map[o1] = j
	h.Map[o2] = i
}
