package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetFiles(path string) (files []string, err error) {
	files = make([]string, 0, 10)
	pathSep := string(".ts")
	suffix := strings.ToUpper(pathSep)
	fileSeperator := string(os.PathSeparator)

	dir, err := ioutil.ReadDir(path)
	checkError(err)

	for _, file := range dir {
		if !file.IsDir() {
			if strings.HasSuffix(strings.ToUpper(file.Name()), suffix) {
				files = append(files, path+fileSeperator+file.Name())
			}
		}
	}

	return files, err
}
