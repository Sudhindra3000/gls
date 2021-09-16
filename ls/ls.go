package ls

import (
	"fmt"
	"log"
	"os"
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
