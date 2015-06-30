// Copyright 2015 Jusong Chen
//
// Author: Jusong Chen (jusong.chen@gmail.com)

package mssql

import (
	"database/sql"
	"fmt"
	"strings"

	//_ "code.google.com/p/odbc"
	_ "github.com/alexbrainman/odbc"
)

//Conn is a wrap up of sql.DB and it is safe to be used concurrently by multiple goroutines.
//field DBName is the sql server session's database context
type Conn struct {
	Host   string
	Login  string
	Passwd string
	DBName string
	DB     *sql.DB
}

//NewCoon creates a connection to MS Sql server
//when empty DBName passed in, Conn.DBName is set to login'sdefault database name
func NewConn(Host, DBName, Login, Passwd string) (*Conn, error) {

	params := map[string]string{
		"driver":   "sql server",
		"server":   Host,
		"database": DBName,
		//"port": 	*msport,
	}

	if len(Login) == 0 {
		params["trusted_connection"] = "yes"
	} else {
		params["uid"] = Login
		params["pwd"] = Passwd
	}

	var connStr string
	for n, v := range params {
		connStr += n + "=" + v + ";"
	}

	// func Open(driverName, dataSourceName string) (*DB, error)
	db, err := sql.Open("odbc", connStr)
	if err != nil {
		return nil, err
	}

	//handling default DB

	var curDBName string
	if err := db.QueryRow("select DB_NAME()").Scan(&curDBName); err != nil {
		return nil, err
	}

	// DBName  specified, and current datbase context is not the same
	if DBName != "" && !strings.EqualFold(DBName, curDBName) {
		return nil, fmt.Errorf("SQL server connection's database context is %s, expecting %s.", curDBName, DBName)
	}

	return &Conn{
		Host:   Host,
		DBName: curDBName, //set database context
		Login:  Login,
		Passwd: Passwd,
		DB:     db,
	}, nil
}

//There
//fixme: this do not always work as once sql.DB has many connection pools
// //
// func setCurDB(db *sql.DB, curDB string) error {

// 	curDBName, err := GetCurDBName(db)
// 	if err != nil {
// 		return err
// 	}

// 	if curDB == curDBName {
// 		return nil
// 	}
// 	//need to set current DB

// 	//set current database
// 	_, err = db.Exec(fmt.Sprintf("use %s", curDB))

// 	curDBName, err = GetCurDBName(db)

// 	if err != nil {
// 		return err
// 	}

// 	if curDB == curDBName {
// 		return nil
// 	} else {
// 		return fmt.Errorf("Set SQL server connection's database context to %s failed", curDB)

// 	}

// }
