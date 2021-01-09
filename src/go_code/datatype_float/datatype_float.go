package main

import(
	"fmt"
)

func main(){
	//浮点数=符号位+指数位+尾数位

	//单精度float32,4字节,-3.403E38~3.403E38
	var num1 float32 = -89.12
	fmt.Println("num1=",num1)

	//双精度float64,8字节,-1.798E308~1.798E308
	var num2 float64 = -89.231231231
	fmt.Println("num2=",num2)

	//尾数部分可能丢失,造成精度损失
	//float64比float32精度高,保存精度高的数用float64
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=",num3,",num4=",num4)

	//浮点默认类型float64
	var num5 = -123.00000901
	fmt.Printf("num5的类型:%T\n",num5)

	//十进制形式
	num6 := 5.12
	num7 :=.123
	fmt.Println("num6=",num6,",num7=",num7)

	//科学技术法
	num8 := 5.1234e2 
	fmt.Println("num8=",num8)
	num9 := 5.1234E2 
	fmt.Println("num9=",num9)
	num10 := 5.1234e-2 
	fmt.Println("num10=",num10)
	num11 := 5.1234E-2 
	fmt.Println("num11=",num11)

}
//浮点型使用细节:
//1:浮点类型有固定的的范围和长度,不收操作系统的影响
//2:默认声明为float64
//3:浮点型常量有两种表示形式:
//   十进制数形式:如5.12,.512(必须由小数点)
//   科学计数法形式:5.1234e2 = 5.12*10^2 , 5.12E-2 = 5.12/10^2
//4:通常使用float64,因为它比float32精度更高