package database

import (
	"database/sql"
	"fmt"

	"github.com/ccxnu/ips-redis-mysql/config"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewMysqlConnection(c *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
