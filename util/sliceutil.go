package util

import "os"

func Filter(slc []os.DirEntry, predicate func(i int) bool) (res []os.DirEntry) {
	for i, v := range slc {
		if predicate(i) {
			res = append(res, v)
		}
	}
	return
}

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
