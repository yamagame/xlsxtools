package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yamagame/xlsxtools"
)

const cmdName = "xlsx2csv"
const cmdShort = "convet xlsx to csv"
const version = "0.1"

var outDirname string

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
		xlsxfile := args[0]
		basename := filepath.Base(xlsxfile)
		if outDirname == "" {
			ext := path.Ext(basename)
			outDirname = filepath.Join(filepath.Dir(xlsxfile), basename[0:len(basename)-len(ext)])
		}
		if _, err := os.Stat(outDirname); os.IsNotExist(err) {
			if err := os.Mkdir(outDirname, 0777); err != nil {
				return err
			}
		}
		f, err := xlsxtools.OpenXLSX(xlsxfile)
		if err != nil {
			return err
		}
		xlsxtools.SaveSheetToCSV(f, outDirname)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&outDirname, "out", "o", "", "output directory")
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
