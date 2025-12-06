package main

import (
	"database/sql"
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

// GetLanguages returns list of available languages from Zotero database
func (zoteroDB *zoteroDBType) GetLanguages() ([]string, error) {
	query := `select distinct upper(ival.value)
	          from items i
	          join itemData idat on idat.itemID = i.itemID and idat.fieldID = 7
	          join itemDataValues ival on ival.valueID = idat.valueID`

	rows, err := zoteroDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []string
	for rows.Next() {
		var language string
		if err := rows.Scan(&language); err != nil {
			return nil, err
		}
		if language != "" { // Filter out empty values
			languages = append(languages, language)
		}
	}

	return languages, nil
}

func (zoteroDB *zoteroDBType) Connect() error {
	// Get user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Build path to Zotero database
	zoteroDB.Path = filepath.Join(homeDir, "Zotero", "zotero.sqlite")
	zoteroDB.buildConnectionString()

	db, err := sql.Open("sqlite", zoteroDB.ConnectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	zoteroDB.DB = db
	return nil
}