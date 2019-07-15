package structure

//Queue 队列
type Queue struct {
	i    int
	data []interface{}
}

//NewQueue 新建一个队列
func NewQueue() *Queue {
	q := &Queue{
		i:    0,
		data: []interface{}{},
	}
	return q
}

//Enqueue 入队
func (q *Queue) Enqueue(val interface{}) bool {
	q.data = append(q.data, val)
	q.i++
	return true
}

//Dequeue 出队
func (q *Queue) Dequeue() interface{} {
	if q.i <= 0 {
		return nil
	}
	res := q.data[0]
	q.data = q.data[1:q.i]
	q.i--
	return res
}

//Length 队列长度
func (q *Queue) Length() int {
	return q.i
}

//IsEmpty 是否为空
func (q *Queue) IsEmpty() bool {
	return q.i == 0
}
