package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gls",
	Short: "golang implementation of the ls command",
	Run:   rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	files, err := ioutil.ReadDir(curDir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Local Flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
