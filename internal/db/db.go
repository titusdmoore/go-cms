package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/titusdmoore/goCms/internal/config"
)

type DB struct {
	Db *sql.DB
}

func InitializeDatabaseConnection(config config.Config) *DB {
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DatabaseName)
	db, err := sql.Open("mysql", connection_string)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DB{
		Db: db,
	}
}

func (db *DB) PurgeDatabase() error {
	tables, err := db.Db.Query("SHOW TABLES;")
	if err != nil {
		return err
	}
	defer tables.Close()

	if _, err := db.Db.Exec("SET FOREIGN_KEY_CHECKS = 0;"); err != nil {
		return err
	}

	for tables.Next() {
		var (
			table_name string
		)

		if err := tables.Scan(&table_name); err != nil {
			return err
		}
		fmt.Printf("Removing Table: %v\n", table_name)

		if _, err := db.Db.Exec(fmt.Sprintf("DROP TABLE %s;", table_name)); err != nil {
			return err
		}
	}

	if _, err := db.Db.Exec("SET FOREIGN_KEY_CHECKS = 1;"); err != nil {
		return err
	}

	return nil
}
