package gunzip

import (
	"io/ioutil"
	"log"
	"os"
)

func FilesInFolder(target string) []os.FileInfo {
	files, err := ioutil.ReadDir(target)
    if err != nil {
        log.Fatal(err)
    }

	return files

}
