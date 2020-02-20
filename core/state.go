package core

import "fmt"

//python的栈结构
type PyStack struct {
	stack       []PyValue //用于存储python的数据类型
	top         int       //当前栈顶位置
	capacity    int       //初始栈的容量
	environment map[string]PyValue
}

func NewStack() *PyStack {
	return &PyStack{
		stack:       make([]PyValue, 10),
		top:         0, //初始栈顶指向0
		environment: make(map[string]PyValue, 0),
		capacity:    10,
	}
}

//pop() 将值推出栈
func (self *PyStack) pop() PyValue {
	//如果栈 元素小于1 则无法出栈
	if self.top < 1 {
		panic("stack underflow")
	}
	//栈顶高度减一
	self.top--
	//栈顶元素
	val := self.stack[self.top]
	//栈顶元素 置nil
	self.stack[self.top] = nil
	//返回pop值
	return val
}

//将值传入栈
func (ps *PyStack) push(val PyValue) {
	//如果栈超出那么栈顶高度那么就报错
	if ps.top > ps.capacity {
		panic("Stack Overflow!")
	}
	ps.stack[ps.top] = val
	ps.top++

}

//检查栈的空闲空间是否还可以 容纳 至少n个值
func (ps *PyStack) check(n int) {
	//总容量 减去 占用的容量
	free := len(ps.stack) - ps.top
	//循环 n - free 次 如果n < free表示空间够用 不需要添加
	for i := free; i < n; i++ {
		ps.stack = append(ps.stack, nil)
	}
}

//将栈顶的值弹出 并打印出来
func (ps *PyStack) print() {
	popval := ps.pop()
	fmt.Println(popval)
}

//将栈顶的2个值弹出再相加
func (ps *PyStack) add() {
	first_val := ps.pop()
	second_val := ps.pop()
	res := Add(first_val, second_val)
	ps.push(res)

}

func (ps *PyStack) STORE_NAME(name PyValue) {
	if name, ok := name.(string); ok {
		val := ps.pop()
		ps.environment[name] = val
	} else {
		panic("Store Name error!")
	}

}

func (ps *PyStack) LOAD_NAME(name PyValue) {
	if name, ok := name.(string); ok {
		val := ps.environment[name]
		ps.push(val)
	} else {
		panic("Load Name error!")
	}

}
