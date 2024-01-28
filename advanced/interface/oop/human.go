package main

import "fmt"

type Human interface {
	eat()
	sleep()
}

type Student struct {
	name string
	age  int
}

func (m *Student) eat() {
	fmt.Println("a man is eating")
}

func (m *Student) sleep() {
	fmt.Println("a man is sleeping")
}

func (m *Student) showName() {
	fmt.Println("name: ", m.name)
}

func (m *Student) showAge() {
	fmt.Println("age: ", m.age)
}

// 构造函数
func NewStudent(name string, age int) *Student {
	st := new(Student)
	st.name = name
	st.age = age
	return st
}

func NewHuman(human Human) {
	human.eat()
}

func main() {
	bob := &Student{
		name: "bob",
		age:  18,
	}

	bob.eat()
	bob.sleep()

	NewHuman(bob)

	tom := NewStudent("Tom", 18)
	tom.showName() // name:  Tom

	tom.showAge() // age:  18

	tom.name = "Jack"
	tom.age = 1

	tom.showName() // name:  Jack
	tom.showAge()  // age:  1
}
