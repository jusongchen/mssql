// Copyright 2015 Jusong Chen
//
// Author: Jusong Chen (jusong.chen@gmail.com)

package mssql

import "database/sql"

// Type Txn implement factory interface
//
type Txn struct {
	conn *Conn
	tx   *sql.Tx
	err_ error
}
