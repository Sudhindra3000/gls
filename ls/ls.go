package ls

import (
	"fmt"
	"gls/util"
	"log"
	"os"
	"sort"
	"strings"
)

func ListFiles(dirPath string, showAll bool) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	listFilesInOrder(entries, showAll)
}

func ListFilesContaining(dirPath string, exp string, showAll bool) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.Contains(entries[i].Name(), exp) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories containing", exp)
		return
	}

	listFilesInOrder(entries, showAll)
}

func ListFilesEndingWith(dirPath string, end string, showAll bool) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.HasSuffix(entries[i].Name(), end) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories ending with", end)
		return
	}

	listFilesInOrder(entries, showAll)
}

func ListFilesStartingWith(dirPath string, start string, showAll bool) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.HasPrefix(entries[i].Name(), start) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories starting with", start)
		return
	}

	listFilesInOrder(entries, showAll)
}

func listFilesInOrder(entries []os.DirEntry, showAll bool) {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
	})
	dirs, files := util.Partition(entries, func(i int) bool { return entries[i].IsDir() })
	for _, entry := range append(dirs, files...) {
		if showAll || !util.HiddenFile(entry.Name()) {
			fmt.Println(entry.Name())
		}
	}
}
