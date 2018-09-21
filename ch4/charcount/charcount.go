package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

//從標準輸入獲得資料,可利用重導向,將檔案的資料匯入
//windows10 	"Get-Content .\testData.txt | go run .\dedup.go" 似乎不能讀取繁體字型
//ubuntu18.04 	"go run .\dedup.go < testData.txt "
//ref https://stackoverflow.com/questions/28504847/powershells-pipe-adds-linefeed/32981504#32981504

func main() {
	dictCounts := make(map[rune]int)
	invaild := 0
	var utflen [utf8.UTFMax + 1]int //不使用index:0 因此需要加1
	input := bufio.NewReader(os.Stdin)
	for {
		r, nbytes, err := input.ReadRune()

		//window10 以ctrl+z 為eof
		//linux 以ctrl+d 為eof
		if err == io.EOF {
			fmt.Println("讀取結束", err)
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "讀取錯誤 %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && nbytes == 1 {
			fmt.Println("解析unicode失敗")
			invaild++
			continue
		}
		dictCounts[r]++  //計算每次字出現的次數
		utflen[nbytes]++ //計算utf8的四種位元組編碼(可變長度的位元組編碼) 所佔的數量
	}

	fmt.Printf("rune \t counts\n")
	for r, v := range dictCounts {
		fmt.Printf("%q \t num=%d \t %[2]x\n", r, v)
	}

	for i, num := range utflen {
		fmt.Printf("type=%d \t num=%d\n", i, num)
	}

	if invaild > 0 {
		fmt.Printf("無法解讀的utf8字元,其數量=%d\n", invaild)
	}
}
