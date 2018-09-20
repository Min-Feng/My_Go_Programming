package main

import (
	"fmt"
	"sort"
)

func main() {

	// * -- 練習 map型別的賦值
	age0 := make(map[string]int)
	age0["Mary"] = 13
	age0["Caeasr"] = 30
	age0["John"] = 28
	fmt.Println("age0=", age0)

	age1 := map[string]int{
		"test": 345,
		"afsf": 396,
	}
	fmt.Println("age1=", age1)

	//只有宣告 沒有使用make給予空間
	//age2為nil,沒有雜湊參考,無法賦值,會產生錯誤
	var age2 map[string]int
	// age2["Mary"] = 13
	// age2["Caeasr"] = 30
	// age2["John"] = 28
	fmt.Println("age2=", age2)

	//注意 age2 age3都為空的map型別
	//但age3有雜湊參考,可以賦值
	//而age2沒有雜湊參考,無法賦值
	age3 := make(map[string]int)
	//age3["Mary"] = 13
	fmt.Println("age3=", age3)
	// * -- 練習 map型別的賦值

	/* 練習無排序的型別 如何排序 */
	//由於已知age的長度,藉由事前分配names空間,避免append的效率低落
	// var names []string
	names := make([]string, 0, len(age0))

	for name := range age0 {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s is %d old.\n", name, age0[name])
	}
	/* 練習無排序的型別 如何排序 */

	/* 判斷map型別的元素是否存在 */
	if age, ok := age0["Mar"]; !ok {
		fmt.Println("此人不存在")
	} else {
		fmt.Println(age)
	}
	/* 判斷map型別的元素是否存在 */

}
