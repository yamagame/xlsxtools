// Pacakage repcpコマンドの構造体
package replace

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// LineReplacer 文字列を入れ替えるインタフェース
type LineReplacer interface {
	Replace(src string) string
}

// File ファイルを管理する構造体
type File struct {
	Filename string
	Replacer LineReplacer
}

// File ファイルを管理する構造体のコンストラクタ
func NewFile(filename string, replacer LineReplacer) *File {
	return &File{
		Filename: filename,
		Replacer: replacer,
	}
}

// MakeDir ディレクトリが存在しなければ作成する
func (x *File) MakeDir() error {
	if !x.IsExist() {
		return os.Mkdir(x.Filename, 0777)
	}
	return nil
}

// IsExist ファイルが存在するか
func (x *File) IsExist() bool {
	_, err := os.Stat(x.Filename)
	return !os.IsNotExist(err)
}

// IsDir ディレクトリかどうか
func (x *File) IsDir() bool {
	fInfo, err := os.Stat(x.Filename)
	if err != nil {
		return false
	}
	return fInfo.IsDir()
}

// Walk ディレクトリに含まれるディレクトリとファイルを返す
func (x *File) Walk(call func(base string) error) error {
	var readDir func(filename, base string) error
	readDir = func(filename, base string) error {
		files, err := os.ReadDir(filepath.Join(filename, base))
		if err != nil {
			return err
		}
		for _, f := range files {
			call(filepath.Join(base, f.Name()))
			if f.IsDir() {
				if err := readDir(x.Filename, filepath.Join(base, f.Name())); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return readDir(x.Filename, "")
}

// Copy 指定した文字を入れ替えながらファイルをコピーする
func (x *File) Copy(dst *File) error {
	if !x.IsExist() {
		return fmt.Errorf("source file %s is not found", x.Filename)
	}
	if !dst.IsExist() {
		return fmt.Errorf("destination directory %s is not found", dst.Filename)
	}
	if !dst.IsDir() {
		return fmt.Errorf("destination %s is not directory", dst.Filename)
	}
	if !x.IsDir() {
		filename := filepath.Base(x.Filename)
		return x.copy(filepath.Join(dst.Filename, filename))
	}
	return x.Walk(func(basefile string) error {
		srcfile := NewFile(filepath.Join(x.Filename, basefile), x.Replacer)
		dstfile := NewFile(filepath.Join(dst.Filename, basefile), x.Replacer)
		if srcfile.IsDir() {
			dstfile.MakeDir()
		} else {
			srcfile.copy(dstfile.Filename)
		}
		return nil
	})
}

func (x *File) copy(filename string) error {
	return x.copyFile(filename)
}

func (x *File) copyFile(filename string) error {
	// 入力ファイルの指定
	srcFile, err := os.Open(x.Filename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 出力ファイルの指定
	dstFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	r := bufio.NewReader(srcFile)
	w := bufio.NewWriter(dstFile)

	for {
		row, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF && len(row) == 0 {
			break
		}

		// 各行を置き換え
		row = x.Replacer.Replace(row)

		_, err = w.WriteString(row)
		if err != nil {
			return err
		}
	}
	return w.Flush()
}
