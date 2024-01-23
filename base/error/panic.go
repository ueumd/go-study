package main

import "fmt"

func fnA() {
	fmt.Println("function fnA")
}

func fnB() {
	panic("function fnB error")
}

func fnB2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("function fnB2")

			fmt.Println(err) // 打印panic 抛出的错误，function fnB2 error
		}
	}()

	panic("function fnB2 error")
}

func fnC() {
	fmt.Println("function fnC")
}
func main() {
	fnA()
	// fnB()
	fnB2()
	fnC()
}

/**
function fnA
panic: function fnB error

goroutine 1 [running]:
main.fnB(...)
        E:/xxx/panic.go:10
main.main()
        E:/xxx/panic.go:18 +0x65
*/

/**
function fnA
function fnB2
function fnB2 error
function fnC
*/
