package xlsxtools

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"

	"github.com/xuri/excelize/v2"
)

func OpenXLSX(filename string) (*excelize.File, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func SaveSheetToCSV(f *excelize.File, outdir string) error {
	sheets := f.GetSheetList()
	digit := len(fmt.Sprintf("%d", len(sheets))) + 1
	form := fmt.Sprintf("%%0%dd-%%s.csv", digit)
	for i, sheet := range sheets {
		records, err := f.GetRows(sheet)
		if err != nil {
			return err
		}
		sheetname := fmt.Sprintf(form, i+1, sheet)
		file, err := os.Create(filepath.Join(outdir, sheetname))
		if err != nil {
			return err
		}
		defer file.Close()
		if err := WriteCSV(file, records); err != nil {
			return err
		}
	}
	return nil
}

func CreateXLSX(filename string, sheetName string, records [][]string) error {
	f := excelize.NewFile()
	sheet := f.GetSheetName(f.GetActiveSheetIndex())
	for y, record := range records {
		for x, value := range record {
			cell, err := excelize.CoordinatesToCellName(x+1, y+1)
			if err != nil {
				return err
			}
			f.SetCellValue(sheet, cell, value)
		}
	}
	return f.SaveAs(filename)
}

func ReadCSV(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	r.FieldsPerRecord = -1
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func WriteCSV(writer io.Writer, records [][]string) error {
	w := csv.NewWriter(writer)
	return w.WriteAll(records)
}

func CompareXLSXWithFilename(source string, target string) (bool, error) {
	s, err := OpenXLSX(source)
	if err != nil {
		return false, err
	}
	t, err := OpenXLSX(target)
	if err != nil {
		return false, err
	}
	result, err := CompareXLSX(s, t)
	if err != nil {
		return false, err
	}
	return result, nil
}

func CompareXLSX(source *excelize.File, target *excelize.File) (bool, error) {
	srcSheets := source.GetSheetList()
	tarSheets := target.GetSheetList()
	if reflect.DeepEqual(srcSheets, tarSheets) {
		for _, sheet := range srcSheets {
			srcCells, err := source.GetRows(sheet)
			if err != nil {
				return false, err
			}
			tarCells, err := target.GetRows(sheet)
			if err != nil {
				return false, err
			}
			if !reflect.DeepEqual(srcCells, tarCells) {
				return false, nil
			}
		}
		return true, nil
	}
	return false, nil
}
