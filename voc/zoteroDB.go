package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
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
	// Get user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Error getting user home directory: %v", err)
		return err
	}

	// Build path to Zotero database
	zoteroDB.Path = filepath.Join(homeDir, "Zotero", "zotero.sqlite")
	zoteroDB.buildConnectionString()

	db, err := sql.Open("sqlite", zoteroDB.ConnectionString)
	if err != nil {
		log.Printf("Error opening database at %s: %v", zoteroDB.Path, err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return err
	}

	zoteroDB.DB = db
	return nil
}