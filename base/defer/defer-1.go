package main

/*
* 匿名返回
 */
func df1() int {
	var x int
	defer func() {
		x++
	}()

	return x
}

/*
*
命名返回
*/
func df2() (x int) {
	defer func() {
		x++
	}()
	return x
}

func main() {
	println(df1()) // 0

	println(df2()) // 1
}
