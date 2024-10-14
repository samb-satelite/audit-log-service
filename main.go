package main

import (
	"fmt"
	"log"
	"os"

	"audit-log/src/infrastructure/api"
	"audit-log/src/infrastructure/database/mysql"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	auditlogs "github.com/samb-satelite/go-audit-logs"
)

// AuditLogModel defines the structure for GORM
type AuditLogModel struct {
	gorm.Model
	Module     string      `json:"module" validate:"required"`
	ActionType string      `json:"action_type" validate:"required"`
	SearchKey  string      `json:"search_key" validate:"required"`
	Before     interface{} `json:"before_data" validate:"required"`
	After      interface{} `json:"after_data" validate:"required"`
	ActionBy   string      `json:"action_by" validate:"required"`
	ActionTime string      `json:"action_time" validate:"required"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file, using environment variables from the host")
	}

	mysql.InitMySQLWriteDB()

	client, err := auditlogs.NewAuditLogClient()
	if err != nil {
		log.Fatalf("Failed to create audit log client: %v", err)
	}
	defer client.Close()

	var consumeError string
	err = client.ConsumeAuditLogs(&consumeError, func(log auditlogs.AuditLog, ack func(bool)) {
		defer ack(true)

		actionTime := log.ActionTime.Format("2006-01-02 15:04:05")

		logData := map[string]interface{}{
			"module":      log.Module,
			"action_type": log.ActionType,
			"search_key":  log.SearchKey,
			"before_data": map[string]string{"before": log.Before},
			"after_data":  map[string]string{"after": log.After},
			"action_by":   log.ActionBy,
			"action_time": actionTime,
		}

		url := os.Getenv("CONSUMER_AUDIT_LOG_URL")

		if err := api.SendDataToAPI(url, logData); err != nil {
			fmt.Printf("Error sending data to API: %v\n", err)
		} else {
			fmt.Println("Module : ", log.Module, "Created successfully!")
		}
	})

	if err != nil {
		log.Fatalf("Failed to start consuming audit logs: %v", err)
	}

	if consumeError != "" {
		log.Fatalf("Error occurred while consuming audit logs: %v", consumeError)
	}

	select {}
}
