package main

import "fmt"

func rev(s []string) {

	//以下為錯誤語法 ,分號之間 只能有一個表達式 expression
	//for i:=,j:=0,len(s)-1 ; i<len(s)/2 ;i++,j-- {

	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []string{"123", "456", "789", "abc", "def", "rest", "dream", "29821667"}
	fmt.Println(s)
	rev(s)
	fmt.Println(s)

	/*
		由於rev函式的參數為 slice的型別
		因此直接 傳入s1會發生錯誤
		應該傳入s1[:]
		slice型別的取得 可使用陣列搭配範圍
		e.g.
		Q_slice=s1[:3]
	*/
	s1 := [4]string{"123", "456", "789", "abc"}
	//rev(s1)
	rev(s1[:])
	fmt.Println(s1)
}
