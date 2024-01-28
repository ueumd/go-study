package i2

import "fmt"

type Phone interface {
	call()
}

type Android struct {
}

func (android *Android) call() {
	fmt.Println("I am Android")
}

type Ios struct {
}

func (ios *Ios) call() {
	fmt.Println("I am Ios")
}

func main() {
	var phone Phone

	phone = new(Android)
	phone.call()

	phone = new(Ios)
	phone.call()
}
