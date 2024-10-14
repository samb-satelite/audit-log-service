package mysql

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbRead *gorm.DB

// InitMySQLReadDB initializes the MySQL database connection.
func InitMySQLReadDB() error {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: Could not load .env file, using environment variables from the host")
	}

	gormConfig := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	}

	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia%%2FJakarta",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))

	connection, err := gorm.Open(mysql.Open(connectionURL), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL database: %w", err)
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

	dbRead = connection
	return nil
}

// GetMySQLReadDB returns the MySQL read database instance.
func GetMySQLReadDB() *gorm.DB {
	return dbRead
}
