// Copyright 2015 Jusong Chen
//
// Author: Jusong Chen (jusong.chen@gmail.com)

package mssql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	//_ "code.google.com/p/odbc"
	_ "github.com/alexbrainman/odbc"
)

func PingMSSQL(c *Conn, count int64, interval time.Duration) {
	var db *sql.DB

	sqlVerCmd := "select @@version"

	now := time.Now()

	fmt.Printf("%v:start connecting to DB\n", now)
	var err error
	err = OpenDB(c)
	if err != nil {
		log.Fatal("connect to DB failed:\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("connect to DB failed:\n", err)
		return
	}
	fmt.Printf("%v:connection to DB established in %v\n", time.Now(), time.Since(now))

	now = time.Now()
	//fmt.Printf("%v:Get DB version\n", now)
	if ver, err := serverVersion(db); err != nil {
		log.Fatal("connect to SQL server failed:\n", err)
		return
	} else {
		log.Printf("connected to SQL server (host=%s, database=%s)\n%s", c.Host, c.DBName, ver)
	}
	fmt.Printf("%v:SQL command %s completed in %v\n", time.Now(), sqlVerCmd, time.Since(now))

	var i int64
	for i = 0; i < count; i = i + 1 {
		time.Sleep(interval)
		now = time.Now()
		if _, err := serverVersion(db); err != nil {
			log.Fatal("connect to SQL server failed:\n", err)
			return
		} else {
			fmt.Printf("%v:%s completed in %v\n", time.Now(), sqlVerCmd, time.Since(now))
		}
	}
}

func serverVersion(db *sql.DB) (sqlVersion string, err error) {
	var v string

	if err = db.QueryRow("select @@version").Scan(&v); err != nil {
		return "", err
	} else {
		return v, nil
	}
}
