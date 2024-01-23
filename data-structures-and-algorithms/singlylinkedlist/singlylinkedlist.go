package singlylinkedlist

import (
	"fmt"
	"strings"
	"sync"
)

/**
有些书籍，把链表做了很细的划分，比如单链表，双链表，循环单链表，循环双链表，其实没有必要强行分类，链表就是从一个数据指向另外一个数据，一种将数据和数据关联起来的结构而已。
1. 单链表，就是链表是单向的，像我们上面这个结构一样，可以一直往下找到下一个数据节点，它只有一个方向，它不能往回找。
2. 双链表，每个节点既可以找到它之前的节点，也可以找到之后的节点，是双向的。
3. 循环链表，就是它一直往下找数据节点，最后回到了自己那个节点，形成了一个回路。循环单链表和循环双链表的区别就是，一个只能一个方向走，一个两个方向都可以走。
*/

// 链表节点
type Node struct {
	value interface{} // 节点元素
	prev  *Node       // 前驱节点
	next  *Node       // 后驱节点
}

// 初始化节点
func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}

type Singly struct {
	size uint64     // 链表元素个数
	head *Node      // 链表头节点指针
	tail *Node      // 链表尾节点指针
	lock sync.Mutex // 并发控制
}

func New(values ...interface{}) *Singly {
	list := &Singly{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *Singly) Add(values ...interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()
	for _, value := range values {
		newNode := &Node{value: value}

		if list.size == 0 {
			// 第一次 前驱节点 后驱节点都指向自已
			list.head = newNode
			list.tail = newNode
		} else {
			list.tail.next = newNode
			list.tail = newNode
		}

		list.size++
	}
}

func (list *Singly) Size() uint64 {
	return list.size
}

func (list *Singly) IsEmpty() bool {
	return list.size == 0
}

func (list *Singly) Clear() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

func (list *Singly) String() string {
	str := "SinglyLinkedList\n"

	values := []string{}

	for head := list.head; head != nil; head = head.next {
		// fmt.Sprintf 同时转换为字符串
		values = append(values, fmt.Sprintf("%v", head.value))
	}

	str += strings.Join(values, ", ")

	return str
}

func (list *Singly) Append(values ...interface{}) {
	list.Add(values...)
}

// 根据index修改value
func (list *Singly) Set(index uint64, value interface{}) {

	if !list.withinRange(index) {
		// 直接添加
		if list.size == index {
			list.Add(value)
			return
		}
	}

	// 其他位置
	curNode := list.head

	for ps := uint64(0); ps != index; {
		ps, curNode = ps+1, curNode.next
	}

	curNode.value = value
}

func (list *Singly) withinRange(index uint64) bool {
	return index >= 0 && index < list.size
}
