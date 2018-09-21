package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, vy := range y {
		if vx, ok := x[k]; !ok || vx != vy {
			//if vy != x[k] {  //若對不存在的元素讀取,則會得到 0 ,可能造成誤判
			return false
		}
	}
	return true
}

func main() {
	x := map[string]int{
		"A": 42,
	}
	y := map[string]int{
		"B": 0,
	}

	isEqual := equal(x, y)
	fmt.Println(isEqual)
}
