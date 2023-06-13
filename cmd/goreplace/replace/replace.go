package replace

import (
	"fmt"
	"strings"
)

// ReplaceString 入れ替える文字列を管理する構造体
type ReplaceString struct {
	Src string
	Dst string
}

// String Package構造体を文字列にする
func (x *ReplaceString) String() string {
	return fmt.Sprintf("{Src: %s, Dst: %s}", x.Src, x.Dst)
}

// ReplaceStrings 複数のPackage構造体管理する
type ReplaceStrings struct {
	Packages []*ReplaceString
}

// NewReplaceStrings 入れ替える文字列を管理する構造体
func NewReplaceStrings(packages []*ReplaceString) *ReplaceStrings {
	return &ReplaceStrings{
		Packages: packages,
	}
}

// Replace 文字列を入れ替える
func (x *ReplaceStrings) Replace(src string) string {
	for _, pkg := range x.Packages {
		src = strings.Replace(src, pkg.Src, pkg.Dst, -1)
	}
	return src
}
