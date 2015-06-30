//Copyright 2015 Jusong Chen
//
//// Author:  Jusong Chen (jusong.chen@gmail.com)

package mssql

import (
	"database/sql"
	"sync"

	log "github.com/golang/glog"
)

type Task interface {
	Process()
	Done()
}

type TaskMaker interface {
	// OpenQuery(*sql.DB) (*sql.Rows, error)
	MakeTask(r *sql.Rows) Task
}

type MakeTaskFunc func(*sql.Rows) Task

//runQueryInParallel take a MakeTaskFunc, a query rowset and process the query rowset in parallel with degree of parallelism of DOP
func runQueryInParallel(DOP int, rs *sql.Rows, f MakeTaskFunc) {

	var wg sync.WaitGroup

	in := make(chan Task)

	wg.Add(1)
	go func() {
		for rs.Next() {
			in <- f(rs)
		}
		err_ := rs.Err()
		if err_ != nil {
			log.Fatal(err_)
		}
		close(in)
		wg.Done()

	}()

	out := make(chan Task)

	for i := 0; i < DOP; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.Process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)

	}()

	for t := range out {
		t.Done()
	}

}

func runQuerySeq(rs *sql.Rows, f MakeTaskFunc) {

	//execute in sequencial order
	for rs.Next() {
		t := f(rs)
		t.Process()
		t.Done()
	}
}

// func RunQuery(e ExecQuery) error {
func ExecQuery(DOP int, rs *sql.Rows, f MakeTaskFunc) {

	if DOP < 1 {
		log.Fatalf("Expect a positive int e.DOP, but got %d", DOP)
	}

	if DOP == 1 {
		runQuerySeq(rs, f)
	} else {
		runQueryInParallel(DOP, rs, f)
	}
}
