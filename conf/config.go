package conf

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // import sqlite driver
	"github.com/spf13/viper"
)

// CaseOne default app dependency objects
type CaseOne struct {
	Env      *viper.Viper
	DBMaster *sqlx.DB
}

// NewConfig set environment values
func NewConfig(env string) *viper.Viper {
	config := viper.New()
	pwd, _ := os.Getwd()

	config.SetConfigName("household")
	config.AddConfigPath(pwd)
	config.SetConfigType("yaml")

	config.SetDefault("Listen", ":8080")

	switch env {
	case "release":
		config.SetDefault("DBMaster", "./main.db")
	case "develop":
		config.SetDefault("DBMaster", "./test.db")
	default:
		config.SetDefault("DBMaster", "./main.db")
	}

	config.ReadInConfig()
	return config
}
func createDB(db *sqlx.DB) {
	tables := map[string]string{
		"user": `
			CREATE TABLE IF NOT EXISTS user(
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				name TEXT UNIQUE, 
				passwd TEXT NOT NULL, 
				create_at NUMERIC NOT NULL DEFAULT (datetime('now'))
			)
		`,
		"category": `
			CREATE TABLE IF NOT EXISTS category(
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				name TEXT UNIQUE,
				budget NUMERIC NOT NULL DEFAULT 0,
				state NUMERIC NOT NULL DEFAULT 0
			)
		`,
		"household": `
			CREATE TABLE IF NOT EXISTS household(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				typeof TEXT NOT NULL,
				category_name TEXT NOT NULL,				
				content TEXT,
				money NUMERIC NOT NULL DEFAULT 0,
				others TEXT,
				author TEXT NOT NULL,
				state NUMERIC NOT NULL DEFAULT 1,				
				create_at NUMERIC NOT NULL DEFAULT (datetime('now')),
				FOREIGN KEY(category_name) REFERENCES category(name),
				FOREIGN KEY(author) REFERENCES user(name)
			)
		`,
	}

	for tableName, query := range tables {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s was created", tableName)
	}
}

// ConnectDB create database connection
func ConnectDB(host string, env *viper.Viper) *sqlx.DB {
	driver := "mysql"
	switch host {
	case "Redshift":
		driver = "postgres"
	case "DBMaster":
		driver = "sqlite3"
	}

	addr := env.GetString(host)
	sess, err := sqlx.Open(driver, addr)
	sess.SetMaxIdleConns(20)
	sess.SetMaxOpenConns(20)

	if err != nil {
		sess.Close()
		log.Fatal(err)
	}
	log.Printf("connect Database: %s\n", host)
	log.Printf("Database driver: %s\n", driver)
	log.Printf("Database address: %s\n", addr)

	if _, err := os.Stat(addr); os.IsNotExist(err) {
		createDB(sess)
	}

	return sess
}

//NewCaseOne create default configuration object
func NewCaseOne(mode string) *CaseOne {
	if mode == "" {
		mode = os.Getenv("APP_ENV")
	}
	env := NewConfig(mode)
	caseOne := &CaseOne{
		Env:      env,
		DBMaster: ConnectDB("DBMaster", env),
	}
	log.Println("configure case: One")

	return caseOne
}
