package main

import (
	"fmt"
)

//每個index欄位 只有1byte大小 因此無號數範圍 0~255
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
// 利用bit shift 將x的bit 右移動,再使用byte()強制轉型 取得x最右邊的8bit
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var x uint64
	x = 999999999999999

	//測試 bit shift 跟 byte()函式
	var i uint
	for i = 0; i < 8; i++ {
		y := x >> (i * 8)
		fmt.Printf("%b \t byte>%b\n", y, byte(y))
	}

}
