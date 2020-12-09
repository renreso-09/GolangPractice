package main

import (
	"fmt" // 標準出力と入力をサポートするライブラリ
	"strings"
	"strconv"
)

func main() {
	var i string
	var s int

	fmt.Scan(&i)

	str := strings.Split(i, "")

	for _, v := range str {
		num,_ := strconv.Atoi(v)
		s += num
	}

	fmt.Printf("%d", s)
}