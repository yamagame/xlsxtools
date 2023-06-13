package pkg

import (
	"fmt"
	"testing"
)

var config = `
pkgs:
  - src: hoge1
    dst: fuge1
  - src: hoge2
    dst: fuge2
`

func TestConfig(t *testing.T) {
	c, _ := ReadConfig(config)
	fmt.Println(c.Pkgs)
}
