package util

func HiddenFile(fileName string) bool {
	return len(fileName) > 0 && fileName[0:1] == "."
}
