package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConPostgres struct{}

func (o *ConPostgres) OpenCon() (*sql.DB, error) {
	env, err := DbEnv.Load(DB_ENGINE_POSTGRES)
	if err != nil {
		return nil, err
	}

	paramconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Host,
		env.Port,
		env.User,
		env.Pass,
		env.Database,
	)

	return sql.Open("postgres", paramconn)
}
