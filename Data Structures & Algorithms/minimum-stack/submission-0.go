type MinStack struct {
	elements []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		mins: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)

	minLen := len(this.mins)
	if minLen == 0 || val <= this.mins[minLen - 1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	elemLen := len(this.elements)
	toRemove := this.elements[elemLen - 1]
	this.elements = this.elements[0:elemLen - 1]

	minLen := len(this.mins)
	if toRemove == this.mins[minLen - 1] {
		this.mins = this.mins[0:minLen - 1]
	}
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements) - 1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins) - 1]
}
