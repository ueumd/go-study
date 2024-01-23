package main

import (
	"fmt"
	"sync"
)

/**
push() 向栈中压入一个元素
pop() 从栈顶取出一个元素
isEmpty() 判断栈是否为空
size() 获取栈中元素的数目
peer() 查询栈顶元素
*/

type Stack struct {
	list []interface{} // 动态数组 切片模拟栈
	size int
	lock sync.Mutex
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) Size() int {
	return stack.size
}

// 入栈
func (stack *Stack) Push(element interface{}) int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.list = append(stack.list, element)

	stack.size += 1

	return stack.size
}

func (stack *Stack) Peak() interface{} {
	if stack.size == 0 {
		return -1
	}
	return stack.list[stack.size-1]
}

// 出栈
func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		return -1
	}

	// 栈顶元素
	value := stack.list[stack.size-1]

	// 切片收缩，但可能占用空间越来越大
	//stack.list = stack.list[0 : stack.size-1]

	// 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	// 开辟内存空间
	// len: stack.size-1 cap: stack.size-1
	newList := make([]interface{}, stack.size-1, stack.size-1)

	// 遍历移动元素
	for i := 0; i < stack.size-1; i++ {
		newList[i] = stack.list[i]
	}

	// 重新赋值
	stack.list = newList

	stack.size -= 1
	return value
}

// 判断一个字符是不是一个运算符 [+, -, *, /]
// 采用ASCII码值进行判断

func (s *Stack) IsOperator(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

// 运算方法
func (s *Stack) Call(operator int) float64 {
	n1 := s.Pop()
	n2 := s.Pop()

	num1 := n1.(float64)
	num2 := n2.(float64)

	res := 0.0

	switch operator {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符有误")

	}
	return res
}

func (s *Stack) Cal(num1, num2, operator int) int {
	res := 0
	switch operator {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符有误")
	}
	return res
}

// 运算优先级设计
func (s *Stack) Priority(oper int) int {
	if oper == 42 || oper == 47 { //[* /] -> 1
		return 1
	} else if oper == 43 || oper == 45 { //[+ -] -> 0
		return 0
	} else {
		return -1
	}

}

func main() {
	stack := &Stack{}
	stack.Push("go")
	stack.Push(1)
	stack.Push("A")

	fmt.Println(stack.list)

	fmt.Println(stack.Pop())

	// fmt.Println(stack.Size())
	// fmt.Println(stack.Peak())
}
