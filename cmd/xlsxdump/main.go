package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

const cmdName = "xlsxdump"
const cmdShort = "dump xlsx files"
const version = "0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + cmdName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version " + version)
	},
}

func getCellBgColor(f *excelize.File, sheet, cell string) string {
	styleID, err := f.GetCellStyle(sheet, cell)
	if err != nil {
		return err.Error()
	}
	fillID := *f.Styles.CellXfs.Xf[styleID].FillID
	fgColor := f.Styles.Fills.Fill[fillID].PatternFill.FgColor
	if fgColor != nil && f.Theme != nil {
		if clrScheme := f.Theme.ThemeElements.ClrScheme; fgColor.Theme != nil {
			if val, ok := map[int]*string{
				0: &clrScheme.Lt1.SysClr.LastClr,
				1: &clrScheme.Dk1.SysClr.LastClr,
				2: clrScheme.Lt2.SrgbClr.Val,
				3: clrScheme.Dk2.SrgbClr.Val,
				4: clrScheme.Accent1.SrgbClr.Val,
				5: clrScheme.Accent2.SrgbClr.Val,
				6: clrScheme.Accent3.SrgbClr.Val,
				7: clrScheme.Accent4.SrgbClr.Val,
				8: clrScheme.Accent5.SrgbClr.Val,
				9: clrScheme.Accent6.SrgbClr.Val,
			}[*fgColor.Theme]; ok && val != nil {
				return strings.TrimPrefix(excelize.ThemeColor(*val, fgColor.Tint), "FF")
			}
		}
		return strings.TrimPrefix(fgColor.RGB, "FF")
	}
	if fgColor != nil {
		return strings.TrimPrefix(fgColor.RGB, "FF")
	}
	return "FFFFFF"
}

var rootCmd = &cobra.Command{
	Use:   cmdName,
	Short: cmdShort,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		f, err := excelize.OpenFile(filename)
		if err != nil {
			return err
		}
		sheets := f.GetSheetList()
		for i, sheet := range sheets {
			fmt.Println(i, sheet)
			rows, err := f.GetRows(sheet)
			if err != nil {
				return err
			}
			for _, c := range f.Styles.Fills.Fill {
				fmt.Println(c.PatternFill.FgColor, c.PatternFill.BgColor)
			}
			for y, cols := range rows {
				for x, cell := range cols {
					fmt.Println(x, y, cell)
					cell, err := excelize.CoordinatesToCellName(x+1, y+1)
					if err != nil {
						return err
					}
					fmt.Println(getCellBgColor(f, sheet, cell))
				}
			}

		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
