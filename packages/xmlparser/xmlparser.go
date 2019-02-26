package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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


func Parse(filename string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened sample.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var omes OMeS
	xml.Unmarshal(byteValue, &omes)


	fmt.Println("#############",omes)

	for i, v := range omes.PMSetup.PMMOResult {
		fmt.Println(i, v.MO.BaseId)
		parseInner(strings.Split(v.NELNBTS.InnerXML, "\n"), 8)
		fmt.Println(i, v.NELNBTS)
	}

}

func parseInner(arrStr []string, idMeasure int64){
	//arrNEL := strings.Join(strings.Split(v.NELNBTS.InnerXML, "\n"), "+")
	nelbs := make(map[string]int, len(arrStr))
	for i, v := range arrStr {
		insideTag := strings.Split(v, "<")
		if len(insideTag) > 1 {
			
			fmt.Println("@!@@@@@@@@@@@@@@@@@@", insideTag[1])
		}

		nelbs[insideTag[0]] = i
	}
	fmt.Println(arrStr, nelbs)
}


