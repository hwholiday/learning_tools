package heap

type Item struct {
	data  interface{}
	ref   int //优先级
	index int //在堆里面的索引
}
type Queue []*Item

func (m Queue) Len() int {
	return len(m)
}

func (m Queue) Less(i, j int) bool {
	return m[i].ref > m[j].ref
}

func (m Queue) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
	m[j].index = j
	m[i].index = i
	//位置交换把对应的索引交换
}

func (m *Queue) Push(d interface{}) {
	d.(*Item).index = len(*m)
	*m = append(*m, d.(*Item))
}

func (m *Queue) Pop() interface{} {
	//抛出后面的一个数据
	l := len(*m)
	s := (*m)[l-1]
	s.index = -1
	*m = (*m)[:l-1]
	return s
}
