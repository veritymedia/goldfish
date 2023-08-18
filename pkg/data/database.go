package data

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"

	"github.com/golang-migrate/migrate/v4/source/iofs"
	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"

	"github.com/veritymedia/goldfish/pkg/config"
	"github.com/veritymedia/goldfish/pkg/utils"
)

//go:embed all:migrations
var schemaFs embed.FS

var (
	dataDir, err = utils.DataDir()
	// dbPath  = dataDir + config.AppName + ".db"
	dbPath  = filepath.Join(dataDir, config.AppName +".db")
	// Pool    *sql.DB
)

func InitialiseDb() (*sql.DB, error) {
	db, err := connectDbPool()
	if err != nil {
		return nil, err
	}

	// TODO: run migrations
	err = runDbMigrations(db)
	if err != nil {
		fmt.Println("Migration error: ", err)
		return nil, err
	}
	fmt.Println("SUCCESS: Migrations completed")
	return db, nil
}

func connectDbPool() (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		fmt.Println("Fatal in sql.Open(path): ", err)
		createDbPath()
		db, err = sql.Open("sqlite3", dbPath)
	}

	if err != nil {
		fmt.Printf("DB 2nd connect retry failed. ")
		return nil, err
	}

	// DbHandler = &Database{Pool: db}
	// fmt.Println("SUCCESS: connectDbPool(), &Pool ", DbHandler.Pool)
	// Pool = db

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR: could not ping DB: ", err)
		createDbPath()
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR: could not ping DB second time: ", err)
		return nil, err

	}
	return db, nil
}

func createDbPath() {
	os.MkdirAll(dataDir, 0755)
	os.Create(dbPath)
}

func runDbMigrations(db *sql.DB) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		fmt.Println("ERROR could not create driver")
		return err
	}
	fmt.Println("SUCCESS created driver")

	d, err := iofs.New(schemaFs, "migrations")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", d, "sqlite3", driver)
	if err != nil {
		return err
	}
	// m, err := migrate.NewWithDatabaseInstance("file:///home/yegor/code/goldfish/pkg/data/migrations/", "sqlite3", driver)

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println("ERROR migration failed: ", err)
		return err
	}

	return nil
}

// func createMigrationTable() {
// 	migration_table := `CREATE TABLE IF NOT EXISTS migrations (
//       id INTEGER PRIMARY KEY,
// 			"name": TEXT NOT NULL,
// 			"up": TEXT NOT NULL,
// 			"down": TEXT NOT NULL,
//       );`
// 	query, err := Pool.Prepare(migration_table)
// 	if err != nil {
// 		fmt.Println("Table could not be created")
// 		log.Fatal(err)
// 	}
// 	query.Exec()
// 	fmt.Println("Table created successfully!")
// }
