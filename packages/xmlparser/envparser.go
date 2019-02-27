package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Data struct {
	XMLName xml.Name	`xml:"data"`
	Mysql struct {
		Host string `xml:"host"`
		User string `xml:"user"`
		Password string `xml:"password"`
	} `xml:"mysql"`
}


func GetEnvironment(filename string) Data {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var data Data
	xml.Unmarshal(byteValue, &data)


	return data

}
