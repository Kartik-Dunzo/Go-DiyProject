package service

import (
	"database/sql"
	"diy_project/config"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = ""
//	dbname   = "pg_database"
//	check_db = "project"
//)
var conf *config.Config

func CheckDatabase() {
	conf = &config.Config_parse
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, "pg_database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("SELECT FROM pg_database WHERE datname = '%s';", conf.Database.DBName)
	fmt.Println(query)
	rows, err := db.Query(query)
	if rows.Next() == false {
		query := fmt.Sprintf("CREATE DATABASE %s;", conf.Database.DBName)
		db.Query(query)
		fmt.Println("Created Database successfully!!!")
		time.Sleep(2 * time.Second)
	}
	rows.Close()
	if err != nil {

		panic(err)
	}
	defer db.Close()

}
