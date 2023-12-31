package connection

import (
	"database/sql"
	"datareporter/config"
	"errors"
)

var AppEnv = config.AppConfig{}
var DbEnv = config.DBConfig{}

const DB_ENGINE_POSTGRES = "postgres"
const DB_ENGINE_SQLITE = "sqlite"

type Generic struct{}

func (cg *Generic) OpenCon(dbEngine string) (*sql.DB, error) {
	switch dbEngine {
	case DB_ENGINE_POSTGRES:
		con := ConPostgres{}
		return con.OpenCon()
	case DB_ENGINE_SQLITE:
		con := ConSqlite{}
		return con.OpenCon()
	default:
		return nil, errors.New("invalid database engine")
	}
}
