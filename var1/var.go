package main

import (
	"fmt"
)

func main() {
	/*
		变量：是一小段内存， 用于存储数据， 在程序运行过程中数值可以改变
		使用：1、申明
			 2、访问，赋值和读取

	*/
	//第一种申明方法：

	var a int
	a = 1
	var b string = "1000"
	fmt.Printf("a is %d,b is %s\n", a, b)

	//第二种申明方法： 类型推断
	var m = 100
	var n = "thisis"
	fmt.Printf("m is : %d,n is %s\n", m, n)

	// 第三种：简短定义
	student := "lixiaohua"
	age := 10
	sex := "male"
	fmt.Printf("student:%s,age:%d,sex:%s", student, age, sex)
	// 多个变量同时定义

	var a1, b1, c1 int = 1, 2, 111
	fmt.Println(a1, b1, c1)
}
