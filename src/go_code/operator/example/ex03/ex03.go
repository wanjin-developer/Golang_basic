package main

import (
	"fmt"
)

//不使用中间变量交换a,b的值

func main(){
	var a int = 100
	var b int = 89
	fmt.Printf("a=%v,b=%v\n",a,b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%v,b=%v\n",a,b)
}