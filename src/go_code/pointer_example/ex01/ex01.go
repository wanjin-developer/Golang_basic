package main

import(
	"fmt"
)
// 获取一个int变量的num的地址,并显示到终端
func main(){
	var i int = 100
	var ptr *int = &i
	fmt.Printf("i value %v,i address %v",i,ptr)

}