package cmd

import (
	"github.com/spf13/cobra"
	"gls/ls"
	"gls/util"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gls",
	Short: "golang implementation of the ls command",
	Run:   rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
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
			ls.ListFilesContaining(dirPath, contains)
			return
		}
		if end, ok := util.EndingWith(args[0]); ok {
			dirPath, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			ls.ListFilesEndingWith(dirPath, end)
			return
		}
		if start, ok := util.StartingWith(args[0]); ok {
			dirPath, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			ls.ListFilesStartingWith(dirPath, start)
			return
		}
		dirPath = args[0]
	default:
		log.Fatalln("Invalid Args")
	}

	ls.ListFiles(dirPath)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
