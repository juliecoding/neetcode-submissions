type entry struct {
	value int
	min int	
}

type MinStack struct {
	entries []entry
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	entry := entry{
		value: val,
	}

	currentMin := this.GetMin()
	if len(this.entries) == 0 || val < currentMin {
		entry.min = val
	} else {
		entry.min = currentMin
	}

	this.entries = append(this.entries, entry)
}

func (this *MinStack) Pop() {
	this.entries = this.entries[0:len(this.entries) - 1]
}

func (this *MinStack) Top() int {
	if (len(this.entries) == 0) {
		return 0
	}
	return this.entries[len(this.entries) - 1].value
}

func (this *MinStack) GetMin() int {
	if (len(this.entries) == 0) {
		return 0
	}
	return this.entries[len(this.entries) - 1].min
}
