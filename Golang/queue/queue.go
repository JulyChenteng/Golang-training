package queue

//A FIFO queue
type Queue []int

//Pushes element into queue
func (queue *Queue) Push(val int) {
	*queue = append(*queue, val)
}

//Pops element form head
func (queue *Queue) Pop() int {
	head := (*queue)[0]
	*queue = (*queue)[1:]

	return head
}

//Returns if the queue is empty or not
func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}
