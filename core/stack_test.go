package core

import (
	"fmt"
	"testing"
	"time"
)

type listx struct {
	in string
	pc int //指令计数器
}

type what_to_execute struct {
	stack        *PyStack
	instructions []listx
	numbers      []int //存储数据
	names        []string
}

func NEwwhat_to_execute() *what_to_execute {
	return &what_to_execute{
		stack: NewStack(),
	}
}
func Newlistx(ins string, num int) {

}
func TestAddestx(t *testing.T) {
	a := NEwwhat_to_execute()
	a.numbers = []int{7, 5, 8}
	a.names = []string{"a", "b"}
	x1 := listx{in: "LOAD_VALUE", pc: 0}
	x2 := listx{in: "LOAD_VALUE", pc: 1}
	x3 := listx{in: "ADD_TWO_VALUES", pc: -1}
	x4 := listx{in: "LOAD_VALUE", pc: 2}
	x5 := listx{in: "ADD_TWO_VALUES", pc: -1}
	x6 := listx{in: "LOAD_VALUE", pc: 1}
	x7 := listx{in: "ADD_TWO_VALUES", pc: -1}
	x8 := listx{in: "STORE_NAME", pc: 0}
	//x9 := listx{in:"PRINT_ANSWER", pc:0}
	a.instructions = append(a.instructions, x1, x2, x3, x4, x5, x6, x7, x8)
	now := time.Now()
	for _, val := range a.instructions {
		instruction, argument := val.in, val.pc
		var args PyValue
		if argument == -1 {
			a.instruction(instruction, nil)
		} else {
			switch instruction {
			case "STORE_NAME":
				args = a.names[argument]
			case "LOAD_NAME":
				args = a.names[argument]
			default:
				args = a.numbers[argument]

			}
			a.instruction(instruction, args)
		}
		fmt.Println("当前python栈状态 :", a.stack.environment)
	}
	spenftime := time.Now().Sub(now)
	fmt.Println(spenftime)
}

func (lx *what_to_execute) instruction(fun string, val PyValue) {
	switch fun {
	case "LOAD_VALUE":
		lx.stack.push(val)
	case "ADD_TWO_VALUES":
		lx.stack.add()
	case "PRINT_ANSWER":
		lx.stack.print()
	case "STORE_NAME":
		lx.stack.STORE_NAME(val)
	case "LOAD_NAME":
		lx.stack.LOAD_NAME(val)

	}
}
