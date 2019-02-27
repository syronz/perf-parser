package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"strconv"

	"github.com/syronz/perf-parser/packages/mongo"
	"github.com/syronz/perf-parser/packages/mysql"
	"github.com/syronz/perf-parser/models"
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
		MeasurementType string `xml:"measurementType,attr"`
		InnerXML string		`xml:",innerxml"`
	}						`xml:"NE-LNBTS_1.0"`
}

type MO struct {
	XMLName xml.Name	`xml:"MO"`
	BaseId string		`xml:"baseId"`
}


type Mvalue struct {
	Tag string
	Amount string
}


func Parse(db mysql.Params, filename string) {

	data := model.Data{}
	data.Date = "32323"


	xmlFile, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}

	arrFilename := strings.Split(filename, ".")[6]

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var omes OMeS
	xml.Unmarshal(byteValue, &omes)


	layout := "2006-01-02T15:04:00.000+00:00:00"
	str :=    omes.PMSetup.StartTime
	t, err := time.Parse(layout, str)

	const createdFormat = "2006-01-02 15:04:05"
	if err != nil {
		fmt.Println(err)
	}

	strDate := fmt.Sprintf(t.Format(createdFormat))

	for _, v := range omes.PMSetup.PMMOResult {
		data.Date = strDate
		data.BaseId = v.MO.BaseId
		data.MeasurementType = v.NELNBTS.MeasurementType
		data.MeasurementTypeID, _ = strconv.Atoi(arrFilename)
		//lastId := db.InsertMeasurement(strDate, v.MO.BaseId, v.NELNBTS.MeasurementType, arrFilename)
		mvalues := parseInner(strings.Split(v.NELNBTS.InnerXML, "\n"), 8)
		data.Mvalues = make(map[string]int, len(mvalues))
		//fmt.Println(">>>>>>>>>>", lastId, strDate, v.MO.BaseId, arrFilename, filename)
		//if lastId > 0 {
			for _, v := range mvalues {
				if v.Tag != "" {
					tmp, errConv := strconv.Atoi(v.Amount)
					if errConv == nil {
						data.Mvalues[v.Tag] = tmp
					}
					//db.InsertMValues(lastId, v.Tag, v.Amount)
					//fmt.Println(lastId, v.Tag, v.Amount)
				}
			}
		//} else {
			//fmt.Println("Duplicate: ", strDate, v.MO.BaseId, arrFilename, filename)
		//}
		mongo.Insert(data)
	}

	os.Remove(filename)

}

// extract the innder XML from NELNBTS's tag
func parseInner(arrStr []string, idMeasure int64) []Mvalue {
	mvalues := make([]Mvalue, len(arrStr))
	nelbs := make(map[string]int, len(arrStr))
	for i, v := range arrStr {
		insideTag := strings.Split(v, "<")
		if len(insideTag) > 1 {
			arrElements := strings.Split(insideTag[1], ">")
			mvalues[i].Tag = arrElements[0]
			mvalues[i].Amount = arrElements[1]
		}

		nelbs[insideTag[0]] = i
	}

	return mvalues
}


