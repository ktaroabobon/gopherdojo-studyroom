package ktaroabobon

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

var d string

func init() {
	flag.StringVar(&d, "d", "img", "変換ファイルの存在するディレクトリ")
}

func extensionExchanger(path, save string) error {
	j, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(j *os.File) {
		_ = j.Close()
	}(j)

	img, err := jpeg.Decode(j)
	if err != nil {
		return err
	}

	p, err := os.Create(save)
	if err != nil {
		return err
	}
	defer func(p *os.File) {
		_ = p.Close()
	}(p)

	err = png.Encode(p, img)
	if err != nil {
		return err
	}
	return nil
}

func kadai1() error {
	flag.Parse()

	err := filepath.Walk(d,
		func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == ".jpg" {
				save := d + filepath.Base(path) + ".png"
				fmt.Println(path)
				err := extensionExchanger(path, save)
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := kadai1()
	if err != nil {
		panic(err)
	}

}
