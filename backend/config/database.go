package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

var config = ConfigDB{}

// Read and parse the configuration file
func (c *ConfigDB) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

// ConnectDB returns initialized sql.DB
func ConnectDB() (*sql.DB, error) {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
