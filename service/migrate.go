package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateDb(db *sql.DB, back *bool, steps *int) {
	driver, derr := mysql.WithInstance(db, &mysql.Config{})
	if derr != nil {
		println(derr.Error())
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"deep-ar-service", driver)
	if err != nil {
		println(err.Error())
		return
	}
	var s int
	if steps != nil {
		s = *steps
	} else {
		s = 0
	}
	if *back {
		err = m.Steps(-s)
	} else {
		if s > 0 {
			err = m.Steps(s)
		} else {
			err = m.Up()
		}
	}
	if err != nil {
		println(err.Error())
		return
	}
}
