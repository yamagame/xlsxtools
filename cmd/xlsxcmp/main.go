package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamagame/xlsxtools"
)

const cmdName = "xlsxcmp"
const cmdShort = "compare xlsx files"
const version = "0.1"

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
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		source := args[0]
		target := args[1]
		s, err := xlsxtools.OpenXLSX(source)
		if err != nil {
			return err
		}
		t, err := xlsxtools.OpenXLSX(target)
		if err != nil {
			return err
		}
		result, err := xlsxtools.CompareXLSX(s, t)
		if err != nil {
			return err
		}
		if result {
			fmt.Println("same")
		} else {
			fmt.Println("different")
			os.Exit(1)
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
