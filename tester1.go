package main

import "fmt"

func AcuSum(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i
	}
	return
}
func RAcuSum(n int) int {
	if n == 1 {
		return 1
	}
	return n + RAcuSum(n-1)
}
func main() {
	res := AcuSum(100)
	fmt.Println(res)
	res1 := RAcuSum(100)
	fmt.Println(res1)
	var fTest FuncType
	fTest = RAcuSum
	result := fTest(100)
	fmt.Println(result)
}

//函数也是一种数据类型 通过type给一个函数类型起名
// 定义FuncType为函数类型
type FuncType func(int) int

func Cal(a, b int, f1 FuncType) (result int) {
	result = f1(a, b)
}

//由f1实现多态
