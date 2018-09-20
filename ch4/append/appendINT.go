package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	/*
		z:=[]int 此為錯誤宣告方法

		正式的變數宣告
		var name type = expression

		短變數宣告
		name := expression

		正式的變數宣告可以省略表達式,只宣告型別
		短變數宣告一定要使用表達式
	*/
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		//x 還有空間可以擴展,根據原始陣列,重新定義slice
		z = x[:zlen]
	} else {
		//若x底層陣列的空間不足夠擴展,則一次擴展兩倍的空間
		//避免每次新增一個元素,都要擴展空間,造成時間的浪費
		zcap := 2 * len(x)
		z = make([]int, zlen, zcap)

		//從x複製所有元素到z,且z指向的底層陣列 與 x已經是不同陣列
		//因為z這個slice是由make重新建構的新空間
		//copy從一個slice複製元素到另一個相同型別的slice !!!
		copy(z, x)
	}

	z[zlen-1] = y
	return z
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	s = appendInt(s, 9)
	fmt.Println(s)
	fmt.Println("")

	//測試copy的功能
	//copy回傳複製的元素數量,通常是兩個slice中較小者的長度
	s1 := []int{9, 8, 7}
	fmt.Println(s1)
	fmt.Println(copy(s1, s), s1)
	fmt.Println(len(s1), cap(s1))
}
