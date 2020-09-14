package main

import (
	"fmt" // 標準出力と入力をサポートするライブラリ
)

func main() {
	// 必要な変数を定義
	var a int
	var b int
	var c int
	var s string

	// 標準入力開始
	fmt.Scan(&a)
	fmt.Scan(&b,&c)
	fmt.Scan(&s)

	// 出力
	fmt.Printf("%d %s", a+b+c, s)
}
