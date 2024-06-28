package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	if len(*q) == 0 {
		return -1
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
