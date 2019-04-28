package main

import (
	"fmt"
	"time"

	"github.com/syronz/perf-parser/packages/gunzip"
	"github.com/syronz/perf-parser/packages/mysql"
	"github.com/syronz/perf-parser/packages/xmlparser"
)

func main() {

	const zipFolder = "./zip-data/"
	const rawFolder = "./raw-data/"
	const envFile = "./environments.xml"

	envs := xmlparser.GetEnvironment(envFile)

	db := mysql.Params{
		Host:     envs.Mysql.Host,
		User:     envs.Mysql.User,
		Password: envs.Mysql.Password,
	}

	db.Connect()
	defer db.Disconnect()

	zipFiles := gunzip.FilesInFolder(zipFolder)
	for _, file := range zipFiles {
		gunzip.Extract(zipFolder+file.Name(), rawFolder)
	}

	rawFiles := gunzip.FilesInFolder(rawFolder)
	start := time.Now()
	var index int
	for _, file := range rawFiles {
		index++
		fmt.Println(time.Since(start), index)
		xmlparser.Parse(db, rawFolder+file.Name())
	}

	fmt.Println("Information added to the database")
	fmt.Println("Duration: ", time.Since(start))

}
