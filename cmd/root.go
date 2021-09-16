package cmd

import (
	"github.com/spf13/cobra"
	"gls/ls"
	"gls/strutil"
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
		if end, ok := strutil.EndingWith(args[0]); ok {
			dirPath, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			ls.ListFilesEndingWith(dirPath, end)
			return
		}
		if start, ok := strutil.StartingWith(args[0]); ok {
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
