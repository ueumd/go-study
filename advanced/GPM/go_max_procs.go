package main

import (
	"fmt"
	"runtime"
)

/*
*
GOMAXPROCS 是 Go 提供的非常重要的一个环境变量。
通过设定 GOMAXPROCS，用户可以调整调度器中 Processor（简称P）的数量。
由于每个系统线程必须要绑定 P ，P 才能把 G 交给 M 执行。

以 P 的数量会很大程度上影响 Go Runtime 的并发表现。
GOMAXPROCS 在 Go 1.5 版本后的默认值是机器的 CPU 核数 （runtime.NumCPU)。
通过下面的代码片段可以获取当前机器的核心数和给 GOMAXPROCS 设置的值。
*/

func main1() {
	fmt.Println("cpus:", runtime.NumCPU())   // cpus: 16
	fmt.Println("goroot:", runtime.GOROOT()) // goroot: D:\go\go
	fmt.Println("archive:", runtime.GOOS)    // archive: windows

}

func main() {
	//设置cpu最大核数
	n := runtime.GOMAXPROCS(1)
	fmt.Println("之前核数=", n)
	for {
		//两个协程抢着输出 0，1；观察01交替密度来观察；核数越大，交替越密
		go fmt.Print(0)
		fmt.Print(1)
	}
}
