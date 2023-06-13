package pkg

import (
	"fmt"
	"strings"
)

type Package struct {
	Src string
	Dst string
}

func (x *Package) String() string {
	return fmt.Sprintf("{Src: %s, Dst: %s}", x.Src, x.Dst)
}

type PackageReplace struct {
	Packages []*Package
}

func NewReplace(packages []*Package) *PackageReplace {
	return &PackageReplace{
		Packages: packages,
	}
}

func (x *PackageReplace) Replace(src string) string {
	for _, pkg := range x.Packages {
		src = strings.Replace(src, pkg.Src, pkg.Dst, -1)
	}
	return src
}
