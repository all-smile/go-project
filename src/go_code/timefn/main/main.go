package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	// fmt.Println("test-str", str)
}

func main() {
	now := time.Now()
	fmt.Printf("time类型:%T, 值： %v \n", now, now)
	fmt.Printf("年：%v \n", now.Year())
	fmt.Printf("月：%v \n", now.Month())
	fmt.Printf("月：%v \n", int(now.Month()))
	fmt.Printf("日：%v \n", now.Day())
	fmt.Printf("时：%v \n", now.Hour())
	fmt.Printf("分：%v \n", now.Minute())
	fmt.Printf("秒：%v \n", now.Second())
	// fmt.Printf("date：%v \n", now.Date())

	// 格式化日期时间
	fmt.Printf("当前年月日： %d-%d-%02d %02d:%02d:%02d \n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	dataStr := fmt.Sprintf("xx当前年月日： %d-%d-%02d %02d:%02d:%02d \n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dataStr类型： %T, 值：%s \n", dataStr, dataStr)

	var curtime = now.Format("2006-01-02 15:04:05")
	fmt.Printf("curtime类型： %T, 值： %s \n", curtime, curtime)
	var time01 = now.Format("2006-01-02") // 年月日
	fmt.Printf("time01类型： %T, 值： %s \n", time01, time01)
	var time02 = now.Format("15:04:05") // 时分秒
	fmt.Printf("time02类型： %T, 值： %s \n", time02, time02)
	var time03 = now.Format("01") // 月份
	fmt.Printf("time03类型： %T, 值： %s \n", time03, time03)

	// 时间常量使用
	i := 0
	for {
		i++
		fmt.Println(i)
		// 休眠
		// time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
		if i >= 20 {
			break
		}
	}

	// unix时间戳和unixnano时间戳
	fmt.Printf("unix时间戳=%v, unixnano时间戳=%v\n", now.Unix(), now.UnixNano())

	// 统计消耗时间
	var startNuix = time.Now().Unix()
	test()
	var endNuix = time.Now().Unix()
	fmt.Println("startNuix = ", startNuix)
	fmt.Println("endNuix = ", endNuix)
	var exectime = endNuix - startNuix
	fmt.Printf("exectime = %v", exectime)
}
