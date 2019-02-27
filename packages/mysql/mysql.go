package mysql

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// connection information saved here
type Params struct {
	Host string
	User string
	Password string
	db *sql.DB
}


func (p *Params) Connect() {
	var err error
    p.db, err = sql.Open("mysql", p.User + ":" + p.Password + "@tcp(" + p.Host + ":3306)/perf_parser")
    if err != nil {
        panic(err.Error())
    }
}

func (p *Params) Disconnect() {
	p.db.Close()
}


func (p *Params) InsertMeasurement(date string, baseID string, measurementType string, measurementTypeId string) int64 {

	// query for insert new row
	stmt, err := p.db.Prepare("INSERT ignore measurement( date, base_id, measurement_type, measurement_type_id) Values(?,?,?,?)")
    if err != nil {
        fmt.Println(err.Error())
    }

	res, err := stmt.Exec(date, baseID, measurementType, measurementTypeId)
	defer stmt.Close()
    if err != nil {
        fmt.Println(err.Error())
    }

	lid, err := res.LastInsertId()
    if err != nil {
        fmt.Println(err.Error())
    }


	return int64(lid)

}

func (p *Params) InsertMValues(measurementID int64, tag string, amount string) {

	// query for insert new row
	stmt, err := p.db.Prepare("INSERT ignore mvalues( id_measurement, tag, amount) Values(?,?,?)")
    if err != nil {
        fmt.Println(err.Error())
    }

	_, err = stmt.Exec(measurementID, tag, amount)
	defer stmt.Close()
    if err != nil {
        fmt.Println(err.Error())
    }

}
