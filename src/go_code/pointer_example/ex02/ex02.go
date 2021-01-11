package main

import(
	"fmt"
)

func main(){
	var i int = 100
	var ptr *int = &i
	fmt.Printf("i value %v ,i address %v\n",i,ptr)
	*ptr = 999
	fmt.Printf("i value %v ,i address %v\n",i,ptr)
	
}