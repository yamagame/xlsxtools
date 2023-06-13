// Pacakage repcpコマンドのメインパッケージ
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamagame/xlsxtools/cmd/repcp/replace"
)

const cmdName = "repcp"
const cmdShort = "copy files while replacing strings"
const version = "0.1"

var configPath string

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
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		srcName := args[0]
		dstName := args[1]
		var srcFile *replace.File
		var dstFile *replace.File
		pkgs := []*replace.ReplaceString{}
		if configPath != "" {
			bytes, err := ioutil.ReadFile(configPath)
			if err != nil {
				panic(err)
			}
			config, err := replace.ReadConfig(string(bytes))
			if err != nil {
				return err
			}
			pkgs = config.Pkgs
		}
		srcFile = replace.NewFile(srcName, replace.NewReplaceStrings(pkgs))
		dstFile = replace.NewFile(dstName, replace.NewReplaceStrings(pkgs))
		return srcFile.Copy(dstFile)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "config file path")
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
