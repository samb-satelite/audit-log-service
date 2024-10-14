package mysql

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbWrite *gorm.DB

// InitMySQLWriteDB initializes the MySQL write database connection.
func InitMySQLWriteDB() error {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: Could not load .env file, using environment variables from the host")
	}

	gormConfig := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	}

	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia%%2FJakarta",
		os.Getenv("MYSQL_WRITE_USER"), os.Getenv("MYSQL_WRITE_PASS"), os.Getenv("MYSQL_WRITE_HOST"), os.Getenv("MYSQL_WRITE_PORT"), os.Getenv("MYSQL_WRITE_DBNAME"))

	connection, err := gorm.Open(mysql.Open(connectionURL), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL write database: %w", err)
	}

	sqlDB, err := connection.DB()
	if err != nil {
		return fmt.Errorf("failed to get DB object: %w", err)
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = 10 // default value
	}
	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = 10 // default value
	}
	connMaxLifetime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))
	if err != nil {
		connMaxLifetime = 300 // default value in seconds
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	dbWrite = connection

	log.Println("Connected to MySQL")
	return nil
}

// GetMySQLWriteDB returns the MySQL write database instance.
func GetMySQLWriteDB() *gorm.DB {
	return dbWrite
}
