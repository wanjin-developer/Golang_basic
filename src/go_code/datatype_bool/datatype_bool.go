package main

import(
	"fmt"
	"unsafe"
)

//布尔类型
func main(){
	//bool类型只允许取true和false
	//只占1个字节
	//适用于逻辑运算,用于程序流程控制,if,for
	var a =false
	fmt.Println("a=",a)
	fmt.Printf("a类型%T,a占用空间%d\n",a,unsafe.Sizeof(a))

}