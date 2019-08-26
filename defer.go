package main

import "fmt"

func main01() {
	defer fmt.Println("aa")
	//defer 延迟调用 main函数结束前调用
	defer fmt.Println("bb")

	fmt.Println("cc")
}

func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println(a, b)
	}(a, b) //a和b在顺序执行的时候已经传参完毕了 只是还没有执行所以出书为10 和 20
	a = 111
	b = 222
	fmt.Println(a, b)
}

//多个defer 先写后出
//且程序崩了也可以运行defer的内容
