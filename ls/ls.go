package ls

import (
	"fmt"
	"gls/util"
	"log"
	"os"
	"sort"
	"strings"
)

func ListFiles(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	listFilesInOrder(entries)
}

func ListFilesEndingWith(dirPath string, end string) {
	var correctEntries []os.DirEntry
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), end) {
			correctEntries = append(correctEntries, entry)
		}
	}

	if len(correctEntries) == 0 {
		fmt.Println("No Files or Directories ending with", end)
		return
	}
	listFilesInOrder(correctEntries)
}

func ListFilesStartingWith(dirPath string, start string) {
	var correctEntries []os.DirEntry
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), start) {
			correctEntries = append(correctEntries, entry)
		}
	}

	if len(correctEntries) == 0 {
		fmt.Println("No Files or Directories starting with", start)
		return
	}
	listFilesInOrder(correctEntries)
}

func listFilesInOrder(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
	})
	dirs, files := util.Partition(entries, func(i int) bool { return entries[i].IsDir() })
	for _, entry := range append(dirs, files...) {
		fmt.Println(entry.Name())
	}
}
