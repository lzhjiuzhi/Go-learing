package main

import "fmt"

func main() {
	a := 10
	str := "mike"
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("name is ", str)
	}
	f1()
	//定义匿名函数 同时调用
	func() {
		fmt.Println("无返回值匿名函数直接调用")
	}()
	//带参数匿名函数
	func(i, j int) {
		fmt.Printf("i=%d,j=%d\n", i, j)
	}(10, 20)
	//有返回值有参数
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = j
			min = i
		} else {
			max = i
			min = j
		}
		return
	}(102, 202)
	println(x, y)
}
