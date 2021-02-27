package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	dbConfigPath *string
	port         *int
	db           *sql.DB
	err          error
)

func main() {
	dbConfigPath = flag.String("dbConfigPath", ".db_config.json", "Path to JSON DB config")
	port = flag.Int("port", 8080, "Port on which miscroservice will run")
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
	db, err = sql.Open("postgres", psqlInfo)
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

	// setup server routes
	setupRoutes()

	// start server
	log.Printf("Start server at :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
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
