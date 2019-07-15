package structure

//Stack 栈结构
type Stack struct {
	i    int
	data []interface{}
}

//NewStack 新建一个栈
func NewStack() *Stack {
	q := &Stack{
		i:    0,
		data: []interface{}{},
	}
	return q
}

//Length 栈长度
func (s *Stack) Length() int {
	return s.i
}

//IsEmpty 是否为空
func (s *Stack) IsEmpty() bool {
	return s.i == 0
}

//Push 入栈
func (s *Stack) Push(k interface{}) {
	s.data = append(s.data, k)
	s.i++
}

//Pop 出栈
func (s *Stack) Pop() (res interface{}) {
	if s.i <= 0 {
		return nil
	}
	s.i--
	res = s.data[s.i]
	s.data = s.data[:s.i]
	return res
}
