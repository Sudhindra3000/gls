package cmd

import (
	"github.com/spf13/cobra"
	"gls/ls"
	"gls/models"
	"gls/util"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gls",
	Short: "golang implementation of the ls command",
	Run:   rootRun,
}
var flags = models.Flags{}

func rootRun(_ *cobra.Command, args []string) {
	var dirPath string
	var err error

	switch len(args) {
	case 0:
		dirPath, err = os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
	case 1:
		if contains, ok := util.Contains(args[0]); ok {
			dirPath, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			ls.ListFilesContaining(dirPath, contains, flags)
			return
		}
		if end, ok := util.EndingWith(args[0]); ok {
			dirPath, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			ls.ListFilesEndingWith(dirPath, end, flags)
			return
		}
		if start, ok := util.StartingWith(args[0]); ok {
			dirPath, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			ls.ListFilesStartingWith(dirPath, start, flags)
			return
		}
		dirPath = args[0]
	default:
		log.Fatalln("Invalid Args")
	}

	ls.ListFiles(dirPath, flags)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Flags
	rootCmd.Flags().BoolVarP(&flags.ShowAll, "all", "a", false, "List all files including hidden files (files with names beginning with a dot)")
	rootCmd.Flags().BoolVarP(&flags.Reverse, "reverse", "r", false, "List the files in the Reverse of the order that they would otherwise have been listed in")
}
