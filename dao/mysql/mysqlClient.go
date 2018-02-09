package mysql

import (
	"fmt"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

const (
	getAllDriverIDSID = "select distinct driver_id,sid from g_credit_result"
)

func InitMySQL(username, password, url, dbname string) *sqlbuilder.Database {
	var settings mysql.ConnectionURL
	settings, err := mysql.ParseURL(username + ":" + password + "@tcp(" + url + ")/" + dbname)

	if err != nil {
		fmt.Println(err)
	}

	var db sqlbuilder.Database
	db, err = mysql.Open(settings)
	if err != nil {
		fmt.Println(err)
	}

	return &db
}
