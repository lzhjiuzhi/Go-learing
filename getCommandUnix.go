package main

import "fmt"
import "os"

func main() {
	//接受用户传递参数 为字符串
	list := os.Args
	n := len(list)
	fmt.Println("n =", n)
}
