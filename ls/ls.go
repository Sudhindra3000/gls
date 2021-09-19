package ls

import (
	"fmt"
	"gls/models"
	"gls/util"
	"log"
	"os"
	"sort"
	"strings"
)

func ListFiles(dirPath string, flags models.Flags) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	listFilesInOrder(entries, flags)
}

func ListFilesContaining(dirPath string, exp string, flags models.Flags) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.Contains(entries[i].Name(), exp) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories containing", exp)
		return
	}

	listFilesInOrder(entries, flags)
}

func ListFilesEndingWith(dirPath string, end string, flags models.Flags) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.HasSuffix(entries[i].Name(), end) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories ending with", end)
		return
	}

	listFilesInOrder(entries, flags)
}

func ListFilesStartingWith(dirPath string, start string, flags models.Flags) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	entries = util.Filter(entries, func(i int) bool { return strings.HasPrefix(entries[i].Name(), start) })
	if len(entries) == 0 {
		fmt.Println("No Files or Directories starting with", start)
		return
	}

	listFilesInOrder(entries, flags)
}

func listFilesInOrder(entries []os.DirEntry, flags models.Flags) {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
	})
	dirs, files := util.Partition(entries, func(i int) bool { return entries[i].IsDir() })
	final := append(dirs, files...)
	if flags.Reverse {
		final = util.Reverse(final)
	}
	for _, entry := range final {
		if flags.ShowAll || !util.HiddenFile(entry.Name()) {
			fmt.Println(entry.Name())
		}
	}
}
