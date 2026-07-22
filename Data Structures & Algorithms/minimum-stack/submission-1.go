type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	stack := make([]int, 0)
	minStack := make([]int, 0)
	return MinStack{
		stack:stack,
		minStack: minStack,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack,val)
	}
}

func (this *MinStack) Pop() {
	v := this.stack[len(this.stack)-1]
	if v == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}