package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	dbConfigPath      *string
)

func main() {
	dbConfigPath = flag.String("dbConfigPath", ".db_config.json", "Path to JSON DB config")
	flag.Parse()

	// load all configs
	loadConfigs()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		loadedDBConfig.Host,
		loadedDBConfig.Port,
		loadedDBConfig.User,
		loadedDBConfig.Password,
		loadedDBConfig.Dbname)

	// connect to DB
	db, err := sql.Open("postgres", psqlInfo)
	errHandle("Could not open DB connection", err)
	defer func() {
		err = db.Close()
		errHandle("Could not close DB connection", err)
	}()
	err = db.Ping()
	errHandle("Could not ping DB", err)

	// load DB select statements
	selectStatements.Statements = make(map[string]string)
	selectStatements.loadStatements()
}

func errHandle(msg string, err error) {
	if err != nil {
		log.Panicln(msg+"\n", err)
	}
}

func loadConfigs() {
	file, err := os.Open(*dbConfigPath)
	errHandle("Could not open provided DB config", err)

	err = loadDBConfig(file)
	errHandle("Could not load provided DB config", err)
}
