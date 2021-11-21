// Package converter 画像ファイルの変換を行うパッケージ
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

// Converter ファイル変換器
type Converter struct {
	/*
		targetDir: 対象ディレクトリ
		fromExt: 変換前拡張子
		toExt: 変換後拡張子
	*/
	targetDir string
	fromExt   string
	toExt     string
}

// Decode 変換前のファイルに対応するデコードを行う
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

// Encode 変換後のファイルに対応するエンコードを行う
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

// Convert ファイルを変換する
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

// Run 変換器のファイル変換を実行する
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

// NewConverter 新しい変換器を生成する
func NewConverter(dir, from, to string) *Converter {
	return &Converter{
		targetDir: dir,
		fromExt:   from,
		toExt:     to,
	}
}
