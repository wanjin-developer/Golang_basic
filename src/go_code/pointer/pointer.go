package main

import (
	"fmt"
)

//Golang指针
func main(){
	//基本数据类型
	var i int = 10
	//i的地址,&i
	//i value 10,i address 0xc0000a2058
	fmt.Printf("i value %v,i address %v\n",i,&i)

	//ptr是一个指针变量
	//ptr的类型是*int
	//ptr本身的值是&i
	var ptr *int = &i
	//ptr type *int,ptr value 0xc0000a2058,ptr address 0xc0000cc020
	fmt.Printf("ptr type %T,ptr value %v,ptr address %v\n",ptr,ptr,&ptr)
	//*ptr取指针指向的值
	fmt.Printf("ptr 指向的值:%v",*ptr)

}

//指针的使用细节:
//1:值类型,都有对应的指针类型,形式为*数据类型,比如int的对应指针就是*int,float32对应的指针类型就是*float
//2:值类型包括:基本数据类型,int系列,float系列,bool,string,数组和结构体struct
//值类型:变量直接存储值,内存通常在栈中分配
//3:引用类型包括:指针,slice切片,map,管道chan,interface接口
//引用类型:变量存储的是一个地址,这个地址对应的空间才真正存储数据(值),内存通常在堆上分配,
//当没有任何变量引用这和地址时,该地址对应的数据空间就成了一个垃圾,由GC来回收
