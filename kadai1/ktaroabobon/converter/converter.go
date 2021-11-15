package converter

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

func Convert(path, save string) error {
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

func Run(d string) error {
	err := filepath.Walk(d,
		func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == ".jpg" {
				save := d + filepath.Base(path) + ".png"
				fmt.Println(path)
				err := Convert(path, save)
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
