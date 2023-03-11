package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yamagame/xlsxtools"
)

const cmdName = "csv2sql"
const cmdShort = "convet csv to sql"
const version = "0.1"

var outFilename string
var indexkey string
var deletekey string
var tablename string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + cmdName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version " + version)
	},
}

func findString(keys []string, val string) int {
	for i, v := range keys {
		if v == val {
			return i
		}
	}
	return -1
}

func insertIntoSQL(table string, keys []string) func(tablename string, indexes, params []string) string {
	return func(tablename string, indexes, params []string) string {
		if table == "" || tablename == table {
			return "INSERT INTO " + tablename + " (" + strings.Join(indexes, ",") + ") VALUES (" + strings.Join(params, ",") + ");"
		}
		return ""
	}
}

func updateSetSQL(table string, keys []string) func(tablename string, indexes, params []string) string {
	return func(tablename string, indexes, params []string) string {
		if table == "" || tablename == table {
			t := make([]string, 0)
			w := make([]string, 0)
			for i, index := range indexes {
				value := params[i]
				if findString(keys, index) >= 0 {
					w = append(w, index+" = "+value)
					continue
				}
				t = append(t, index+" = "+value)
			}
			return "UPDATE " + tablename + " SET " + strings.Join(t, ",") + " WHERE " + strings.Join(w, " AND ") + ";"
		}
		return ""
	}
}

func deleteSQL(table string, keys []string) func(tablename string, indexes, params []string) string {
	return func(tablename string, indexes, params []string) string {
		if table == "" || tablename == table {
			w := make([]string, 0)
			for i, index := range indexes {
				value := params[i]
				if findString(keys, index) >= 0 {
					w = append(w, index+" = "+value)
					continue
				}
			}
			return "DELETE FROM " + tablename + " WHERE " + strings.Join(w, " AND ") + ";"
		}
		return ""
	}
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
		indexkeys := strings.Split(indexkey, ",")
		deletekeys := strings.Split(deletekey, ",")
		gensql := insertIntoSQL(tablename, indexkeys)
		if indexkey != "" {
			gensql = updateSetSQL(tablename, indexkeys)
		} else if deletekey != "" {
			gensql = deleteSQL(tablename, deletekeys)
		}
		sqls := xlsxtools.CreateSQL(records, gensql)
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
	rootCmd.Flags().StringVarP(&indexkey, "key", "k", "", "where index key")
	rootCmd.Flags().StringVarP(&tablename, "table", "t", "", "sql table name")
	rootCmd.Flags().StringVarP(&deletekey, "del", "d", "", "delete index key")
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
