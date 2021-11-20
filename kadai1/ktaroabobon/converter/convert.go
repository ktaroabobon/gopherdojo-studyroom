package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Converter struct {
	targetDir string
	fromExt   string
	toExt     string
}

func (c Converter) Decode(r io.Reader) (image.Image, error) {
	switch c.fromExt {
	case "jpg", "jpeg":
		return jpeg.Decode(r)
	case "png":
		return png.Decode(r)
	case "gif":
		return gif.Decode(r)
	}
	return nil, nil
}
func (c Converter) Encode(w io.Writer, m image.Image) error {
	switch c.toExt {
	case "jpg", "jpeg":
		return jpeg.Encode(w, m, &jpeg.Options{Quality: 100})
	case "png":
		return png.Encode(w, m)
	case "gif":
		return gif.Encode(w, m, &gif.Options{NumColors: 256})
	}
	return nil
}

func (c Converter) Convert(path, save string) error {
	o, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(o *os.File) {
		_ = o.Close()
	}(o)

	img, err := c.Decode(o)
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

	err = c.Encode(p, img)
	if err != nil {
		return err
	}
	return nil
}

func (c Converter) Run() error {
	err := filepath.Walk(c.targetDir,
		func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == "."+c.fromExt {
				dir, file := filepath.Split(path)
				save := filepath.Join(dir, file[:len(file)-len(filepath.Ext(file))]+"."+c.toExt)
				fmt.Println(path)
				fmt.Println(save)
				err = c.Convert(path, save)
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

func NewConverter(dir, from, to string) *Converter {
	return &Converter{
		targetDir: dir,
		fromExt:   from,
		toExt:     to,
	}
}
