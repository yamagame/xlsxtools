package xlsxtools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/xlsxtools"
)

func TestCompareXLSX(t *testing.T) {
	comp := func(a string, b string) bool {
		r, err := xlsxtools.CompareXLSXWithFilename(a, b)
		assert.NoError(t, err)
		return r
	}
	test1 := "./testdata/test1.xlsx"
	test2 := "./testdata/test2.xlsx"
	assert.Equal(t, true, comp(test1, test1))
	assert.Equal(t, true, comp(test2, test2))
	assert.Equal(t, false, comp(test1, test2))
	assert.Equal(t, false, comp(test2, test1))
}
