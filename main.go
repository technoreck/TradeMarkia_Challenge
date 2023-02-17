package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"

	"github.com/technoreck/TradeMarkia_Challenge/datatype"
)

func main() {

	file, err := os.Open("tt230101.xml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	// Read the XML data
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var xml_data datatype.XmlFile
	err = xml.Unmarshal(data, &xml_data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	json_data, err := json.MarshalIndent(xml_data, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = os.WriteFile("DB_Result.json", json_data, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := sql.Open("postgres", "postgres://irrdzyqn:rav3LEPxmu-lOBam54ikWtOdiM7BZ3QC@tiny.db.elephantsql.com/irrdzyqn")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}

}
