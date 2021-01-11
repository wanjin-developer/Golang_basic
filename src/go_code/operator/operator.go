package main

import (
	"fmt"
)

//运算符
func main(){
	//算数运算符,对数值类型变量进行运算
	//+(正号) -(负号) +(加) -(减) *(乘) /(除) %(模除) ++(自加) --(自减) +(字符串相加)
	
	//如果运算符都是整数,那么除后,去掉小数部分,保留整数部分
	//2
	fmt.Println(10/4)
	var n1 float32 =10/4
	//2
	fmt.Println(n1)

	//如果我们希望保留小数部分,则需要有浮点数参与运算
	var n2 float32 = 10.0/4
	//2.5
	fmt.Println(n2)

	//取模的本质:a%b=a-a/b*b
	//1
	fmt.Println("10%3=",10 % 3)
	//-1
	fmt.Println("-10%3=",-10 % 3)
	//1
	fmt.Println("10%-3=",10 % -3)
	//-1
	fmt.Println("-10%-3=",-10 % -3)

	var n3 int = 10
	n3++
	//11
	fmt.Println("n3=",n3)
	n3--
	//10
	fmt.Println("n3=",n3)

	//算数运算符的细节:
	//1:对于/,它的整数和小数除是有区别的,整数之间做除法时,只保留整数部分而舍弃小数部分
	//2:当对一个数取模时,可以等价a%b=a-a/b*b
	//3:Golang的自增自减只能当做一个独立语言使用,不能这样使用,b:=a++ 或者 b:=a--
	//4:Golang的++,--只能写在变量的后面,不能写在变量的前面,只有a++,a--,没有++a,--a
	//5:Golang去掉c/java中的自增自减的容易混淆的写法,让Golang更简洁,统一

	//赋值运算符
	//将某个变量运算后的值,赋给指定的变量
	//=:讲一个表达式的值赋给左值
	//+=:相加后赋值
	//-=:相减后赋值
	//*=:相乘后赋值
	///=:相除后赋值
	//%=:求余后赋值
	//二进制赋值运算符:
	//<<=:左移后赋值
	//>>=:右移后赋值
	//&=:按位与后赋值
	//^=:按位异或后赋值
	//!=:按位或后赋值
	var i int
	//基本赋值 
	i = 10
	//10
	fmt.Println("i=",i)

	a := 9
	b := 2
	//定义一个临时变量
	t := a
	a = b
	b = t
	//a=2,b=9
	fmt.Printf("a=%v,b=%v\n",a,b)

	a += 17
	//a=19
	fmt.Printf("a=%v\n",a)
	a -=10
	//a=9
	fmt.Printf("a=%v\n",a)
	a *= 2
	//a=18
	fmt.Printf("a=%v\n",a)
	a /= 3
	//a=6
	fmt.Printf("a=%v\n",a)
	a %= 2
	//a=0
	fmt.Printf("a=%v\n",a)

	var c int 
	c = a + 3
	//赋值运算符的执行顺序是从右向左
	//c=3
	fmt.Printf("c=%v\n",c)

	//赋值运算符左边只能是变量,右边可以是变量,表达式,常量值
	//任何有值都可以看做表达式
	var d int

	d = a
	//d=0
	fmt.Printf("d=%v\n",d)

	d = 8 + 2 * 8
	//d=24
	fmt.Printf("d=%v\n",d)

	d = 890
	//d=890
	fmt.Printf("d=%v\n",d)


	//比较运算符/关系运算符
	//== != < > <= >=
	//关系运算符的结果都是bool型,要么是true,要么是false
	//关系表达式经常用在if结构的条件中或循环结构的条件中
	//关系运算符组成的表达式称为关系表达式:a > b
	//==不能写成=
	ba := 1
	bb := 2
	//false
	fmt.Println(ba == bb)
	//true
	fmt.Println(ba != bb)
	//true
	fmt.Println(ba < bb)
	//false
	fmt.Println(ba > bb)
	//true
	fmt.Println(ba <= bb)
	//false
	fmt.Println(ba >= bb)

	//逻辑运算符
	//用于连接多个条件(一般来讲就是关系表达式),最终结果是一个bool值
	//&&:逻辑与,两边为真true,则为true,否则为false
	//||:逻辑或,两边有一个true,则为true,否则为false
	//!:逻辑非,如果条件为true,则逻辑为false,否则为true.取反

	var age int = 40
	//ok1
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age <40 {
		fmt.Println("ok2")
	}

	//ok1 ok2
	if age > 30 || age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 || age <40 {
		fmt.Println("ok2")
	}

	//ok1
	if age > 30 {
		fmt.Println("ok1")
	}
	if !(age > 30){
		fmt.Println("ok2")
	}
	//逻辑运算符注意细节:
	//&&也叫短路与,如果第一个条件为false,则第二个条件不会判断,最终结果false
	//||也叫短路或,如果第一个条件为true,则第二个条件不会判断,最终结果为true
	var li int = 10
	if li > 9 && test() {
		fmt.Println("短路与ok1")
	}
	if li < 9 || test() {
		fmt.Println("短路或ok2")
	}

	//位运算符
	//其他运算符:&取地址,*指针变量

	//运算符的优先级
	//1:只有单目运算符,赋值运算符是从右向左运算的
	//后缀>单目>乘除>加减>移位>关系>相等关系>
	//       按位AND>按位XOR>按位OR>逻辑AND>逻辑OR>赋值运算符>逗号
	//大致顺序:
	//1:括号,++,--
	//2:单目运算
	//3:算数运算符
	//4:移位
	//5:关系运算
	//6:位运算
	//7:逻辑运算
	//8:赋值运算
	//9:逗号

}

//声明一个函数(测试)
func test() bool {
	fmt.Println("test...")
	return true
}