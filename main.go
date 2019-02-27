package main

import (
	"fmt"

	"github.com/syronz/perf-parser/packages/gunzip"
	"github.com/syronz/perf-parser/packages/xmlparser"
	"github.com/syronz/perf-parser/packages/mysql"
)


func main() {
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
	//defer db.Disconnect()

	//lastId := db.InsertMeasurement("2019-02-23 12:00:00", "NE-MRBTS-20", "LTE_S1AP", "9000")
	//fmt.Println(lastId)


	zipFiles := gunzip.FilesInFolder(zipFolder)
	for _, file := range zipFiles {
		gunzip.Extract(zipFolder + file.Name(), rawFolder)
	}


	rawFiles := gunzip.FilesInFolder(rawFolder)
	for _, file := range rawFiles {
		xmlparser.Parse(db, rawFolder + file.Name())
	}


	fmt.Println("Information added to the database")
	//xmlparser.Parse("./test/sample02.xml")
	//xmlparser.Parse(db, "./raw-data/PM.LTEOMS-1.OSS.20190226.124500.8038")


}
