package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/datatools/mysql"
)

func main() {

	config, err := mysql.NewDBConfig("", "", "", false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	fmt.Printf("Connection established on: %s\n", config.Config())
	// err := mysql.Check()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// Check performs a connection check on the mysql database connection
func Check() error {
	dbconfig, err := mysql.NewDBConfig("", "", "", false)
	if err != nil {
		log.Errorln(err)
	}

	err = dbconfig.Save(".mysql.cfg")
	if err != nil {
		log.Errorln(err)
	}

	db, err := mysql.NewDbConnect(dbconfig, "")
	if err != nil {
		log.Errorln(err)
	}

	// defer the close until  the main function is done
	// This is not normal in a running application. It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
	defer db.Close()

	// perform query test
	response, err := db.Query("SHOW DATABASES;")
	if err != nil {
		log.Errorln(err)
	}

	log.Info("MySQL query response: ", response)
	return nil

}
