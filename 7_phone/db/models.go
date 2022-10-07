package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// todo fill out

var user = "postgres"
var password = "postgres"
var db = "testing"
var host = "localhost"
var port = "5432"
var ssl = "disable"
var timezone = "Etc/UTC"
var dbConn *gorm.DB

type NumberID struct {
	ID string `uri:"id" binding:"required"`
}

type Number struct {
	ID          string `gorm:"primaryKey" json:"id"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (x *Number) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}

func GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, db, port, ssl, timezone)
}

func Test(s string) {
	fmt.Println(s)
}
