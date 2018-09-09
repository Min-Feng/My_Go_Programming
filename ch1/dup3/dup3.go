package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filePath := range os.Args[1:] {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		} else {
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}
	for key, val := range counts {
		fmt.Printf("%s\t%d\n", key, val)
	}
}
