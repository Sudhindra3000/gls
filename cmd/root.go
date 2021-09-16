package cmd

import (
	"github.com/spf13/cobra"
	"gls/ls"
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
	// Local Flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
