package database

import (
	"database/sql"
	"fmt"
	"projet-forum/utils"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() (*sql.DB, error) {
	user := utils.GetEnvWithDefault("DB_USER", "")
	pwd := utils.GetEnvWithDefault("DB_PWD", "")
	host := utils.GetEnvWithDefault("DB_HOST", "")
	port := utils.GetEnvWithDefault("DB_PORT", "")
	name := utils.GetEnvWithDefault("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pwd, host, port, name)
	dbContext, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := dbContext.Ping(); err != nil {
		dbContext.Close()
		return nil, err
	}

	return dbContext, nil
}
