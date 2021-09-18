package util

import "os"

func Partition(slc []os.DirEntry, predicate func(i int) bool) (p1 []os.DirEntry, p2 []os.DirEntry) {
	for i, v := range slc {
		if predicate(i) {
			p1 = append(p1, v)
		} else {
			p2 = append(p2, v)
		}
	}
	return
}
