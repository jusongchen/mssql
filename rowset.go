// Copyright 2015 Jusong Chen
//
// Author: Jusong Chen (jusong.chen@gmail.com)

package mssql

import "database/sql"

// Type Rowset implement factory interface
//
type Rowset struct {
	conn    *Conn
	sqlText string
	args_   []interface{}
	rows    *sql.Rows
	err     error
}

func OpenRowSet(conn *Conn, sqlText string, args ...interface{}) (*Rowset, error) {
	var args_ []interface{}
	for _, arg := range args {
		args_ = append(args_, arg)
	}

	rs := &Rowset{conn: conn, sqlText: sqlText, args_: args_}
	rs.rows, rs.err = rs.conn.db.Query(rs.sqlText, rs.args_...)
	return rs, rs.err
}

func (s *Rowset) Next() bool {
	return s.rows.Next()
}

func (s *Rowset) Close() error {
	return s.rows.Close()
}

func (r *Rowset) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}
