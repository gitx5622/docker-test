package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required. 
)

const (
	hostname = "localhost"
	hostport = 5432
	username = "postgres"
	password = 1165
	databasename = "docker-test"
	)

// MigrateDb migrates scripts present in /db/migrations(on container) during app startup	
func MigrateDb() error {
	migrationDir := "db/migrations"
	pgconstring := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
	 hostname, hostport, username, password, databasename)
	log.Printf("dir is %v",migrationDir)
	log.Printf("pg conn string is %v",pgconstring)

	db,err := sql.Open("postgres",pgconstring)

	if err != nil {
		log.Printf("Unable to connect to postgre db. Error is %v",err)
	}

	// defer the closing part
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping postgres db. Error is %v",err)
	}

	driver,err := postgres.WithInstance(db,&postgres.Config{})

	if err != nil {
		log.Fatalf("could not start database. Error is  %v",err)
	}
	// make sure you import the file dependency. If not file:// won't work
	mPath := fmt.Sprintf("file://%s", migrationDir)
	log.Printf("path of migration directory is %s",mPath)
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationDir),"postgres",driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated")
	return nil
}