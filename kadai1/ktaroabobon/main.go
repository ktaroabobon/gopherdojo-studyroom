package main

import (
	"flag"
	"gituhub.com/ktaroabobon/gopherdojo-studyroom/kadai1/ktaroabobon/converter"
)

var d string

func init() {
	flag.StringVar(&d, "d", "img", "変換ファイルの存在するディレクトリ")
}
func main() {
	flag.Parse()

	e := converter.Run(d)
	if e != nil {
		panic(e)
	}
}
