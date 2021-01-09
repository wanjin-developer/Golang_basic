package main

//import "fmt"
//import "unsafe"
import (
	"fmt"
	"unsafe"
)

//go数据类型
//基本数据类型: 数值型 int,int8,int16,int32,int64,uint无符号,uint8,uint16,uint32,uint64,byte，
//             浮点型 float32,float64，
//             字符型 byte保存单个字符，
//             布尔型 bool，
//             字符串 string
//派生/复杂数据类型：指针 Pointer，数组，结构体struct，管道Channel，函数，切片slice，接口interface，map
func main(){
	//有符号整数类型
	// int 32位系统占4个字节,64位系统占8个字节
	var n1 int = 1
	fmt.Println("n1=",n1)

	//int8范围：-128~127，1字节，超过范围会报错：overflows int8
	var n2 int8 = 127
	fmt.Println("n2=",n2)

	//int16范围：-2^15~2^15-1，2字节
	var n3 int16 = 2222
	fmt.Println("n3=",n3)

	//int32范围：-2^31~2^31-1，4字节
	var n4 int32 = 2222222
	fmt.Println("n4=",n4)

	//int64范围：-2^63~2^63-1，8字节
	var n5 int64 = 2222222222
	fmt.Println("n5=",n5)

	//无符号整数类型
	//uint 32位系统占4个字节,64位系统占8个字节
	var un0 uint  = 1
	fmt.Println(un0)

	//uint8，1字节，0~255
	var un1 uint8 = 255
	fmt.Println("un1=",un1)

	//uint16,2字节,0~2^16-1
	var un2 uint16 = 25555
	fmt.Println(un2)

	//uint32,4字节,0~2^32-1
	var un3 uint32 = 2555555555
	fmt.Println(un3)

	//uint64,8字节,0~2^64-1
	var un4 uint64 = 255555555555555555
	fmt.Println(un4)

	//rune,有符号,与int32一样,区别表示一个Unicode码
	var rn1 rune = 255555555
	fmt.Println(rn1)

	//byte,无符号,与uint8等价,存储字符时选用byte
	var bn1 byte = 255
	fmt.Println(bn1)

	//查看变量类型%T
	var t1 int32 = 255555555
	fmt.Printf("t1的类型:%T\n",t1)

	//如何在程序中查看某个变量的占用字节大小和数据类型
	var t2 int64 = 10
	//unsafe.Sizeof(t2) 是unsafe包的一个函数,可以返回t2所占的字节数
	fmt.Printf("t2 的类型:%T ,t2占用的字节数:%d \n",t2,unsafe.Sizeof(t2))

}
//整数的使用细节:
//1:有符号,无符号,表示范围不同,int和uint字节数与系统有关
//2:整型默认int型
//3:查看某个变量的子节点大小和数据类型
//4:遵循保小不保大原则,即在程序正确的前提下尽量使用占用空间小的数据类型[年龄使用uint8/byte]
//5:bit,指的是计算机中最小存储单位.byte,指的是计算机中基本存储单元. 1byte=8bit
