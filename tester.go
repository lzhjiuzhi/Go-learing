package main

import "fmt"

func main() {
	Myfunc1()
	Myfunc2(123, "caonima")
	Myfunc3(1)
	Myfunc3(1, 2)
	Myfunc3(1, 2, 3)
	shit := Myfunc4()
	fmt.Println(shit)
	ddasd := Myfunc5()
	fmt.Println(ddasd)

	max, min := Myfunc6(11, 22)
	fmt.Printf("max = %d\n,min = %d\n", max, min)
}

//无参数无返回值
func Myfunc1() {
	a := 666
	fmt.Println("a=", a)
}
func Myfunc2(a int, b string) {
	fmt.Println("a =", a)
	fmt.Printf(b)
}

// 不定参数类型 ...int 不定参数只能放在最后一个参数

func Myfunc3(args ...int) {
	fmt.Println("len(args) = ", len(args))
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

// 一个返回值
func Myfunc4() int {
	return 444123
}

// 给返回值名字
func Myfunc5() (xx int) {
	xx = 444
	return
}

// 可以有多个返回值
func Myfunc6(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		min = a
		max = b
	}
	return
}
