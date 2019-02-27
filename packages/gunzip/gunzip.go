package gunzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func Extract(filename string, rawFolder string){

	if filename == "" {
		fmt.Println("Usage : gunzip sourcefile.gz")
		os.Exit(1)
	}

	gzipfile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader, err := gzip.NewReader(gzipfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer reader.Close()

	fileArr := strings.Split(filename, "/")

	newfilename := rawFolder + strings.TrimSuffix(fileArr[len(fileArr) -1], ".gz")

	writer, err := os.Create(newfilename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Remove(filename)

	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

