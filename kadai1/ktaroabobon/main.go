package main

import (
	"flag"

	"gituhub.com/ktaroabobon/gopherdojo-studyroom/kadai1/ktaroabobon/converter"
)

var d, from, to string

func init() {
	flag.StringVar(&d, "d", "img", "変換ファイルの存在するディレクトリ")
	flag.StringVar(&from, "f", "jpg", "変換前のファイル拡張子")
	flag.StringVar(&to, "t", "png", "変換後のファイル拡張子")
}
func main() {
	flag.Parse()

	c := converter.NewConverter(d, from, to)
	e := c.Run()
	if e != nil {
		panic(e)
	}
}
