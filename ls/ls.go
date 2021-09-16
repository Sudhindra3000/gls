package ls

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ListFiles(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func ListFilesEndingWith(dirPath string, end string) {
	var correctEntries []string
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, end) {
			correctEntries = append(correctEntries, name)
		}
	}

	if len(correctEntries) == 0 {
		fmt.Println("No Files or Directories ending with", end)
		return
	}
	for _, name := range correctEntries {
		fmt.Println(name)
	}
}

func ListFilesStartingWith(dirPath string, start string) {
	var correctEntries []string
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, start) {
			correctEntries = append(correctEntries, name)
		}
	}

	if len(correctEntries) == 0 {
		fmt.Println("No Files or Directories starting with", start)
		return
	}
	for _, name := range correctEntries {
		fmt.Println(name)
	}
}
