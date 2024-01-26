package _defer

//老虞要学GoLang-函数(上) http://www.cnblogs.com/howDo/archive/2013/06/04/GoLang-function.html
import (
	"fmt"
	"strings"
)

/**
func (p myType ) funcName ( a, b int , c string ) ( r , s int ) {
    return
}

关键字——func
方法名——funcName
入参——— a,b int,b string
返回值—— r,s int
函数体—— {}

*/

/**
Go通过给函数标明所属类型，来给该类型定义方法，
 p myType 即表示给myType声明了一个方法， p myType 不是必须的。
 如果没有，则纯粹是一个函数，通过包名称访问。packageName.funcationName
*/

// 定义新的类型double，主要目的是给float64类型扩充方法
type double float64

// 判断a是否等于b 给double类型增加IsEqual方法
func (a double) IsEqual(b double) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001
}

// 判断a是否等于b 纯粹方法
func IsEqual(a, b float64) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001
}

func test() {
	var a double = 1.999999
	var b double = 1.9999998
	fmt.Println(a.IsEqual(b))
	fmt.Println(a.IsEqual(3))
	fmt.Println(IsEqual((float64)(a), (float64)(b)))
}

// 可变参数
func add(a int, b int, arg ...int) (sum int) {
	sum = a + b
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

// 返回值
func ReadFull(r Reader, buf []byte) (n int, err error) { //返回值 n err
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	/**
	支持给返回值命名后，实际上就是省略了变量的声明，return时无需写成return n,err 而是将直接将值返回
	*/
	return
}

// 函数赋值
var fc = func(msg string) {
	fmt.Println("you say :", msg)
}

/*
*
defer 延迟函数
defer 延迟执行，在声明时不会立即执行，而是在函数return后时按照后进先出的原则依次执行每一个defer。
这样带来的好处是，能确保我们定义的函数能百分之百能够被执行到，这样就能做很多我们想做的事，如释放资源，清理数据，记录日志等

结论：
defer 是在执行完return 后执行
defer 先进后出
触发异常也会走defer语句。
*/
func deferFunc() int {
	index := 0

	fc := func() {
		fmt.Println(index, "匿名函数1")
		index++

		defer func() {
			fmt.Println(index, "匿名函数1-1")
			index++
		}()
	}

	defer func() {
		fmt.Println(index, "匿名函数2")
		index++
	}()

	defer fc()

	return func() int {
		fmt.Println(index, "匿名函数3")
		index++
		return index
	}()
}

func testDEfer() int {
	defer func() {
		fmt.Println("匿名函数2")
	}()

	//
	return func() int {
		fmt.Println("匿名函数1")
		return 1
	}()
}

func defer_call() {
	//先进后出
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常") //触发异常也会走defer语句。

	/**
	打印后
	打印中
	触发异常
	打印前
	*/
}

func calc(a, b int) (sum int, agv int) {
	sum = a + b
	agv = (a + b) / 2
	return
}

// 闭包 可以存储到变量中作为参数传递给其它函数，能够被函数动态的创建和返回。
func Adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func init() {
	// _标识符，用来忽略返回值
	res, _ := calc(100, 200)
	fmt.Println("res=", res)

	sum := add(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("sum=", sum)

	//函数赋值
	fmt.Printf("%T \n", fc) //func(string)
	fc("hello,my love")     //you say : hello,my love

	//直接执行
	func(msg string) {
		fmt.Println("say: ", msg) //say : I love to code
	}("I love to code")

	deferFunc()
	/*
		0 匿名函数3
		1 匿名函数1
		2 匿名函数1-1
		3 匿名函数2
	*/

	testDEfer()
	/*
		匿名函数1
		匿名函数2
	*/

	// defer_call()

	f := Adder()
	fmt.Println(f(1))  // 1
	fmt.Println(f(10)) //11
	fmt.Println(f(20)) //31

	f1 := makeSuffix(".jpg")
	fmt.Println(f1("c"))   //c.jpg
	fmt.Println(f1("c++")) //c++.jpg
}
