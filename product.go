package main

import (
	"fmt" // 標準出力と入力をサポートするライブラリ
)

func main() {
	var a int
	var b int
	
	fmt.Scan(&a, &b)

	if a*b%2 == 0 {
		fmt.Printf("Even")
	} else {
		fmt.Printf("Odd")
	}
}