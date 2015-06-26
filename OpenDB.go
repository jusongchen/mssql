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

type Conn struct {
	Host   string
	Login  string
	Passwd string
	DBName string
	db     *sql.DB
}

func openDB(c *Conn) error {

	params := map[string]string{
		"driver":   "sql server",
		"server":   c.Host,
		"database": c.DBName,
		//"port": 	*msport,
	}

	if len(c.Login) == 0 {
		params["trusted_connection"] = "yes"
	} else {
		params["uid"] = c.Login
		params["pwd"] = c.Passwd
	}

	var conn string
	for n, v := range params {
		conn += n + "=" + v + ";"
	}

	var err error
	c.db, err = sql.Open("odbc", conn)
	return err

}

// //DBName return the connection's database context
// func (c *Conn) GetDBName() string {
// 	return c.DBName
// }

//OpenDB open a MS SQL server database to use.
//When c.DBName is empty, the login's default database is open and c.DBName is set to login's default database after OpenDB()
func OpenDB(c *Conn) error {

	if err := openDB(c); err != nil {
		return err
	}
	//handling default DB

	var curDBName string
	if err := c.db.QueryRow("select DB_NAME()").Scan(&curDBName); err != nil {
		return err
	}

	//DBName not
	if c.DBName == "" {
		c.DBName = curDBName
		return nil
	}

	if !strings.EqualFold(c.DBName, curDBName) {
		return fmt.Errorf("SQL server connection's database context is %s, expecting %s.", curDBName, c.DBName)
	}

	return nil

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
