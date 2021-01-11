package main

import (
	"fmt"
)

//离放假还有97天,问还有多少个星期和几天?
func main(){
	var weeks int 
	var days int
	weeks = 97 / 7
	days = 97 % 7
	//离放假还有13周6天
	fmt.Printf("离放假还有%v周%v天",weeks,days)

}