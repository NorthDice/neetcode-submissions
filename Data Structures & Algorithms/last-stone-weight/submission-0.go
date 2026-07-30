type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i,j int) bool {
	return h[i] > h[j]
} 
func (h minHeap) Swap (i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *minHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *minHeap) Pop () any {
	n := len(*h)
	x:= (*h)[n-1]
	*h = (*h)[:n-1]
	return x
} 

func Init(slice []int) *minHeap {
	h := &minHeap{}
	heap.Init(h)

	for _,num := range slice {
		heap.Push(h,num)
	}
	return h
}
func lastStoneWeight(stones []int) int {
	h:=Init(stones)
//stones=[2,3,6,2,4] 6 4 3 2 2 /  2 2 1/  2 2
	for h.Len() > 1 {
		stone1 := heap.Pop(h)
		if h.Len() == 0 {
			return stone1.(int)
		}
		stone2 := heap.Pop(h)
		if stone1 != stone2 {
			diff := stone1.(int)-stone2.(int)
			heap.Push(h,diff)	
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}
