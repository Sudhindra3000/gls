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
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.HasSuffix(entries[i].Name(), end) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories ending with", end)
		return
	}

	listFilesInOrder(entries)
}

func ListFilesStartingWith(dirPath string, start string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.HasPrefix(entries[i].Name(), start) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories starting with", start)
		return
	}

	listFilesInOrder(entries)
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
