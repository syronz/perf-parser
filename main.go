package main

import (
	"time"
	"fmt"

	"github.com/syronz/perf-parser/packages/gunzip"
	"github.com/syronz/perf-parser/packages/xmlparser"
	"github.com/syronz/perf-parser/packages/mysql"
)


func main() {
	start := time.Now()
	const zipFolder = "./zip-data/"
	const rawFolder = "./raw-data/"
	const envFile	= "./environments.xml"

	envs := xmlparser.GetEnvironment(envFile)

	db := mysql.Params{
		Host: envs.Mysql.Host,
		User: envs.Mysql.User,
		Password: envs.Mysql.Password,
	}

	db.Connect()
	defer db.Disconnect()

	zipFiles := gunzip.FilesInFolder(zipFolder)
	for _, file := range zipFiles {
		gunzip.Extract(zipFolder + file.Name(), rawFolder)
	}


	rawFiles := gunzip.FilesInFolder(rawFolder)
	for _, file := range rawFiles {
		xmlparser.Parse(db, rawFolder + file.Name())
	}


	fmt.Println("Information added to the database")
	fmt.Println("Duration: ", time.Since(start))


}
