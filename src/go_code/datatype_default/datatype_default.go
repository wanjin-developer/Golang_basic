package main

import(
	"fmt"
)

//数据类型默认值,零值
func main(){
	//整型->0
	//浮点型->0
	//字符串->""
	//布尔类型->false
	var a int
	var b float32
	var c float64
	var isMarryied bool
	var name string
	//%v表示按照变量的值输出
	fmt.Printf("a=%v,b=%v,c=%v,isMarryied=%v,name=%v\n",a,b,c,isMarryied,name)

}