package structure

//HeapInt int类型作为元素的堆
type HeapInt []int

//Len 返回数据长度
func (h HeapInt) Len() int { return len(h) }

//Less 比较两个数值
func (h HeapInt) Less(i, j int) bool {
	return h[i] < h[j]
}

//Seap 交换两个数值
func (h HeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//Pop 返回一个值
func (h *HeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Push
func (h *HeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
