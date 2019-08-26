package main

import "fmt"

func main() {

	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	m3 := make(map[int]string, 10)
	m3[1] = "mike"
	fmt.Println("m3 = ", m3)
	m4 := map[int]string{1: "mike", 2: "nimabi", 3: "nmasile"}
	m4[1] = "shabi"
	m4[2] = "sandy"
	fmt.Println("m4 = ", m4)
	for key, value := range m4 {
		fmt.Printf("%d======> %s/n", key, value)
	}
	//判断key是否存在 第一个为key对应value 第二个为ok true 不在村false
	_, ok := m4[1]
	if ok == true {
		fmt.Println("OK")
	} else {
		fmt.Println("NIL")
	}
	//删除某一个key
	delete(m4, 1) //删除key为1的内容
	fmt.Println(m4)
}
