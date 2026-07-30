type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := maxHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		stone1 := heap.Pop(&h).(int)
		stone2 := heap.Pop(&h).(int)
		if stone1 != stone2 {
			heap.Push(&h, stone1-stone2)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return h[0]
}