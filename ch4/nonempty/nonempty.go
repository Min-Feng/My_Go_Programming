package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	//測試 nonempty 是否改變原先的陣列
	testString := [4]string{"one", "", "three", "test"}
	data := testString[:]

	fmt.Printf("%q\n", testString)
	data = nonempty(data)
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", testString)

	fmt.Println("")

	//測試 nonempty2 是否改變原先的陣列
	testString2 := [4]string{"one", "", "three", "test"}
	data2 := testString2[:]

	fmt.Printf("%q\n", testString2)
	data2 = nonempty2(data2)
	fmt.Printf("%q\n", data2)
	fmt.Printf("%q\n", testString2)
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			//迴圈內append的次數一定小於cap(strings)
			//因此不怕append的行為,參考了不同的底層陣列
			out = append(out, s)
		}
	}

	out = append(out, "avd") //直到此步驟,都是參考同樣的底層陣列testString2,同時更改了陣列中的元素

	/*
		此步驟開始,append加入元素後,超過容許值cap(strings2)
		因此對於out來說,底層陣列的參考改變了,不再是strings2
	*/
	out = append(out, "ewef")
	out = append(out, "kfjgji")
	return out
}

//!-alt
