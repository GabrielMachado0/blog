package main

import (
	"database/sql"
	"fmt"
	dbConfig "blog_adaptive/dbconfig"
	
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type Accounts struct {
	user_id  int
	name 	 string
	password string
	email 	 string
}

func main() {
	fmt.Printf("Acessando %s...", dbConfig.DbName)

	db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSoucerName)

	if err != nil {
		panic(err.Error())
	}else {
		fmt.Println("Connected!")
	}

	sqlStatement, err := db.Query("SELECT * FROM " + dbConfig.TableName)
	checkErr(err)

	for sqlStatement.Next() {
		var acc Accounts

		sqlStatement := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", dbConfig.TableName)

		err = db.QueryRow(sqlStatement, 1).Scan(&acc.user_id, &acc.name, &acc.password, &acc.email)
		checkErr(err)

		fmt.Printf("%s\t%s \n", acc.email, acc.password)
	}

	defer db.Close()
}

/*func sqlSelect() {

	sqlStatement, err := db.Query("SELECT name, password, email FROM " + dbConfig.TableName)
	checkErr(err)

	for sqlStatement.Next() {
		var acc Accounts

		sqlStatement := fmt.Sprintf("SELECT name, password, email FROM %s WHERE user_id = $1", dbConfig.TableName)

		err = db.QueryRow(sqlStatement, 1).Scan(&acc.name, &acc.password, &acc.email)
		checkErr(err)

		fmt.Printf("%d\t%s\t%s \n", acc.name, acc.password, acc.email)
	}
}*/