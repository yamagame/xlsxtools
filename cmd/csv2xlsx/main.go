package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/yamagame/xlsxtools"
)

const cmdName = "csv2xlsx"
const cmdShort = "convet csv to xlsx"
const version = "0.1"

var sheetName string
var outFilename string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + cmdName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version " + version)
	},
}

var rootCmd = &cobra.Command{
	Use:   cmdName,
	Short: cmdShort,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		csvfile := args[0]
		var filename string
		if outFilename != "" {
			// オプションのファイル名を採用
			filename = outFilename
		} else {
			// 拡張子をcsvに入れ替え
			ext := path.Ext(csvfile)
			filename = csvfile[0:len(csvfile)-len(ext)] + ".xlsx"
		}
		r, err := os.Open(csvfile)
		if err != nil {
			return err
		}
		records, err := xlsxtools.ReadCSV(r)
		if err != nil {
			return err
		}
		return xlsxtools.CreateXLSX(filename, sheetName, records)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&sheetName, "sheet", "s", "Sheet1", "sheet name")
	rootCmd.Flags().StringVarP(&outFilename, "out", "o", "", "output xlsx filename")
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
