package main

import (
	"fmt"

	"github.com/syronz/perf-parser/packages/gunzip"
	"github.com/syronz/perf-parser/packages/xmlparser"
)


func main() {
	const zipFolder = "./zip-data/"
	const rawFolder = "./raw-data/"

	zipFiles := gunzip.FilesInFolder(zipFolder)

	for _, file := range zipFiles {
		fmt.Println(file.Name())
		gunzip.Extract(zipFolder + file.Name(), rawFolder)
	}



	xmlparser.Parse("./test/sample02.xml")


}
