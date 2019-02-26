package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

)

type OMeS struct {
	XMLName xml.Name	`xml:"OMeS"`
	PMSetup PMSetup			`xml:"PMSetup"`
}

type PMSetup struct {
	XMLName xml.Name		`xml:"PMSetup"`
	StartTime string		`xml:"startTime,attr"`
	PMMOResult []PMMOResult	`xml:"PMMOResult"`
}

type PMMOResult struct {
	XMLName xml.Name	`xml:"PMMOResult"`
	MO	MO			`xml:"MO"`
	NELNBTS struct {
		InnerXML string		`xml:",innerxml"`
	}						`xml:"NE-LNBTS_1.0"`
}

type MO struct {
	XMLName xml.Name	`xml:"MO"`
	BaseId string		`xml:"baseId"`
}

/*
type User struct {
	Users   []User   `xml:"user"`
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}
*/


func Parse() {
		// Open our xmlFile
	xmlFile, err := os.Open("./test/sample02.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened sample.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var omes OMeS
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &omes)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	fmt.Println("#############",omes)

	for i, v := range omes.PMSetup.PMMOResult {
		fmt.Println(i, v.MO.BaseId)
		fmt.Println(i, v.NELNBTS)
	}

	/*
	for i := 0; i < len(datar.Users.Users); i++ {
		fmt.Println("User Type: " + omes.PMSetup)
		//fmt.Println("User Name: " + users.Users[i].Name)
		//fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		//fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)
	}
	*/
}


