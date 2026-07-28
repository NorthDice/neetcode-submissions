
type minHeap []int

func(h minHeap) Len() int {
    return len(h)
}
func(h minHeap) Less(i,j int) bool {
    return h[i] < h[j]
}
func(h minHeap) Swap(i,j int) {
    h[i],h[j] = h[j],h[i]
}
func(h *minHeap) Push(x any) {
    *h = append(*h,x.(int))
}
func(h *minHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
type KthLargest struct {
  minH *minHeap
  k int  
}


func Constructor(k int, nums []int) KthLargest {
    h := &minHeap{}
    heap.Init(h)

    obj := KthLargest{
        minH: h,
        k: k,
    }

    for _, num := range nums {
        obj.Add(num)
    }
    return obj
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.minH,val)

    if this.minH.Len() > this.k {
        heap.Pop(this.minH)
    }

    return (*this.minH)[0]
}
