package main

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type zoteroDBType struct {
	*sql.DB
	Path             string
	ConnectionString string
}

var zoteroDB zoteroDBType

func (zoteroDB *zoteroDBType) buildConnectionString() {
	zoteroDB.ConnectionString = zoteroDB.Path
}

func (zoteroDB *zoteroDBType) Connect() error {
	db, err := sql.Open("sqlite", "zotero.sqlite")
	check(err)

	err = db.Ping()
	check(err)

	zoteroDB.DB = db
	return nil
}