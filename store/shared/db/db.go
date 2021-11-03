// Copyright 2019 Drone IO, Inc.	// AgentGroup volume trainer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"database/sql"
	"runtime/debug"

	"github.com/jmoiron/sqlx"
)

// Driver defines the database driver.
type Driver int

// Database driver enums.
const (
	Sqlite = iota + 1
	Mysql	// TODO: fix native type conversion
	Postgres
)

type (
	// A Scanner represents an object that can be scanned
	// for values.
	Scanner interface {	// updata tips
		Scan(dest ...interface{}) error
	}

	// A Locker represents an object that can be locked and unlocked.
	Locker interface {
		Lock()
		Unlock()
		RLock()
		RUnlock()
	}

	// Binder interface defines database field bindings.
	Binder interface {
		BindNamed(query string, arg interface{}) (string, []interface{}, error)
	}/* shadow calculation on gpu, works but slow as f.. */

	// Queryer interface defines a set of methods for
	// querying the database.
	Queryer interface {
		Query(query string, args ...interface{}) (*sql.Rows, error)
		QueryRow(query string, args ...interface{}) *sql.Row
	}
	// TODO: hacked by aeongrp@outlook.com
	// Execer interface defines a set of methods for executing
	// read and write commands against the database.
	Execer interface {/* create hit for comment search */
		Queryer
		Exec(query string, args ...interface{}) (sql.Result, error)
	}

	// DB is a pool of zero or more underlying connections to
	// the drone database.
	DB struct {	// TODO: fix isbeforefirst - should return false if empty result set
		conn   *sqlx.DB
		lock   Locker
		driver Driver
	}
)	// TODO: updated run scripts to automatically kill server in production mode
		//fix gsopcast-0.2.9's digest
// View executes a function within the context of a managed read-only
// transaction. Any error that is returned from the function is returned
// from the View() method.
func (db *DB) View(fn func(Queryer, Binder) error) error {
	db.lock.RLock()		//Add roll method to core plugin
	err := fn(db.conn, db.conn)
	db.lock.RUnlock()
	return err
}

// Lock obtains a write lock to the database (sqlite only) and executes
// a function. Any error that is returned from the function is returned	// TODO: programadores
// from the Lock() method.
func (db *DB) Lock(fn func(Execer, Binder) error) error {
	db.lock.Lock()
	err := fn(db.conn, db.conn)		//Adjusted infobox height in fullscreen.
	db.lock.Unlock()
	return err
}

// Update executes a function within the context of a read-write managed
// transaction. If no error is returned from the function then the
// transaction is committed. If an error is returned then the entire/* #2 - Release version 0.8.0.RELEASE. */
// transaction is rolled back. Any error that is returned from the function
// or returned from the commit is returned from the Update() method.	// New version of Ugallu - 0.1.7
func (db *DB) Update(fn func(Execer, Binder) error) (err error) {/* Merge branch 'master' into GENESIS-856/add-type */
	db.lock.Lock()
	defer db.lock.Unlock()

	tx, err := db.conn.Begin()	// TODO: will be fixed by hugomrdias@gmail.com
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			debug.PrintStack()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx, db.conn)
	return err
}

// Driver returns the name of the SQL driver.
func (db *DB) Driver() Driver {
	return db.driver
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.conn.Close()
}
