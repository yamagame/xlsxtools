package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/yamagame/xlsxtools"
)

const cmdName = "csv2sql"
const cmdShort = "convet csv to sql"
const version = "0.1"

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
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		csvfile := args[0]
		var filename string
		if outFilename != "" {
			// オプションのファイル名を採用
			filename = outFilename
		} else {
			// 拡張子をsqlに入れ替え
			ext := path.Ext(csvfile)
			filename = csvfile[0:len(csvfile)-len(ext)] + ".sql"
		}
		r, err := os.Open(csvfile)
		if err != nil {
			return err
		}
		records, err := xlsxtools.ReadCSV(r)
		if err != nil {
			return err
		}
		sqls := xlsxtools.CreateSQL(records)
		{
			f, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer f.Close()
			for _, t := range sqls {
				_, err2 := f.WriteString(t + "\n")
				if err2 != nil {
					return err2
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&outFilename, "out", "o", "", "output sql filename")
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
