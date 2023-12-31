package connection

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type ConSqlite struct{}

func (o *ConSqlite) OpenCon() (*sql.DB, error) {
	env, err := DbEnv.Load(DB_ENGINE_SQLITE)
	if err != nil {
		return nil, err
	}

	path, err := os.Getwd()
	if err != nil {
		return nil, errors.New("current location unknow")
	}

	dabaseFile := path + "/../db/sqlite/" + env.Database
	if _, err := os.Stat(dabaseFile); err != nil {
		return nil, fmt.Errorf("sqlite database %v not found", env.Database)
	}

	return sql.Open("sqlite3", dabaseFile)
}
