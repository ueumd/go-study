package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)

	// 1e6 = 1*10^6 = 100000

	//时间戳（秒）：1636616300;
	//时间戳（纳秒）：1636616300046251000;
	//时间戳（毫秒）：1636616300046;
	//时间戳（纳秒转换为秒）：1636616300;
}

func timeTest() {
	now := time.Now()
	log.Println("时间戳（秒）：", now.Unix())       // 输出：时间戳（秒） ： 1698897501
	log.Println("时间戳（毫秒）：", now.UnixMilli()) // 输出：时间戳（毫秒）： 1698897501523
	log.Println("时间戳（微秒）：", now.UnixMicro()) // 输出：时间戳（微秒）： 1698897501523778
	log.Println("时间戳（纳秒）：", now.UnixNano())  // 输出：时间戳（纳秒）： 1698897501523778100

	// 时间戳格式化
	fmt.Println(time.Unix(time.Now().Unix(), 0).Format(time.DateTime))        // 秒 2023-11-02 12:03:12
	fmt.Println(time.UnixMilli(time.Now().UnixMilli()).Format(time.DateTime)) // 毫秒 2023-11-02 12:00:08

	// 2023-11-02 12:00:08

	timeStr := "2023-11-02 12:09:33" // 1698898173102
	stamp, _ := time.ParseInLocation(time.DateTime, timeStr, time.Local)
	fmt.Println(stamp)        // 2023-11-02 12:09:33 +0800 CS
	fmt.Println(stamp.Unix()) // 1698898173 秒

}

func time2() {
	fmt.Println(time.Hour)
	fmt.Println(time.Second)
	fmt.Println(time.Now().Hour())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2021-11-11 15:32:09
	fmt.Println(time.Now().Format(time.UnixDate))         // Thu Nov 11 15:32:09 CST 2021

	// 秒 1636616165
	fmt.Println(time.Now().Unix())

	// 获取指定日期的时间戳
	dt, _ := time.Parse("2006-01-02 15:04:05", "2021-11-11 15:32:09")

	// 1636644729
	fmt.Println(dt.Unix())

	// 1636702329
	fmt.Println(time.Date(2021, 11, 12, 15, 32, 9, 0, time.Local).Unix())
}

func main() {
	timeTest()
}
