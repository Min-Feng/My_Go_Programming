package main

import "fmt"

type flags uint

const (
	FlagUp flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPoinrToPoint
	FlagMulticast
)

//IsUp 查詢：利用常數Flag 配合bit運算 判斷Flag是否開啟
func IsUp(v flags) bool { return v&FlagUp == FlagUp }

//TurnDown 關閉：設置Flag為down 利用位元清除運算 &^
func TurnDown(v *flags) { *v &^= FlagUp }

//SetBroadcast 設置：flag 應該用 ｜運算  而不是＆
//|運算 類似位元相加的概念
func SetBroadcast(v *flags) { *v |= FlagBroadcast }

//IsCast 對比第一種判斷Flag是否開啟
//若有多重條件判斷 先使用｜運算 把需要判斷的flag相加(｜運算在位元運算有類似相加的概念) 然後使用＆ 判斷是否為零
func IsCast(v flags) bool { return v&(FlagMulticast|FlagBroadcast) != 0 }

func main() {
	var v flags = FlagBroadcast | FlagUp
	fmt.Printf("%08b %t \n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%08b %t \n", v, IsUp(v))
}
