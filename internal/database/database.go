package database

// import (
// 	"database/sql"
// 	"log"
// 	"os"

// 	_ "github.com/jackc/pgx/v5/stdlib" // Import pgx as a database/sql driver
// 	"gorm.io/gorm"
// )

// type DB struct {
// 	SQL *gorm.DB
// }

// var dbConn *DB

// // ConnectSQL initializes and returns a database connection
// func ConnectSQL(dsn string) (*DB, error) {
// 	// Use provided DSN or fallback to environment variable
// 	if dsn == "" {
// 		dsn = os.Getenv("DATABASE_URL")
// 	}
// 	if dsn == "" {
// 		log.Fatalln("DATABASE_URL environment variable is not set")
// 	}

// 	// Open a database connection
// 	sqlDB, err := sql.Open("pgx", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Test the connection
// 	err = sqlDB.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Assign the database connection to dbConn
// 	dbConn = &DB{SQL: sqlDB}
// 	return dbConn, nil
// }

// // GetDB provides access to the current DB instance
// func GetDB() *sql.DB {
// 	if dbConn == nil || dbConn.SQL == nil {
// 		log.Fatalln("Database connection is not initialized")
// 	}
// 	return dbConn.SQL
// }
