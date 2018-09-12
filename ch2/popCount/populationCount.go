package main

import (
	"fmt"
)

var pc [256]byte

// func init (){
// 	for i:=0 ; i<len(pc) ; i++{
// 		pc[i]=pc[i/2]+byte(i)
// 	}
// }

func init (){
	for i:= range pc{
		pc[i]=pc[i/2]+byte(i)
	}
}

func main(){
	for i := 0; i<len(pc) ;i++{
		fmt.Println(pc[i])
	}
}