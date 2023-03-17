package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)

	time.Sleep(time.Second * 5)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("fail: db.AutoMigrate, %v\n", err)
	}
	if err := db.AutoMigrate(&Reason{}); err != nil {
		log.Fatalf("fail: db.AutoMigrate, %v\n", err)
	}
	if err := db.AutoMigrate(&MentalPoint{}); err != nil {
		log.Fatalf("fail: db.AutoMigrate, %v\n", err)
	}
	if err := db.AutoMigrate(&ReasonsOnMentalPoints{}); err != nil {
		log.Fatalf("fail: db.AutoMigrate, %v\n", err)
	}

	sampleUsers := Users{
		{Id: "1", Name: "A", Email: "a@gmail.com", UId: "1"},
		{Id: "2", Name: "B", Email: "b@gmail.com", UId: "2"},
		{Id: "3", Name: "C", Email: "c@gmail.com", UId: "3"},
	}

	if err := sampleUsers.CreateUsers(); err != nil {
		log.Fatalf("fail: db.Create Users, %v\n", err)
	}

	sampleReasons := Reasons{
		{Id: "1", Reason: "A", UserId: "1"},
		{Id: "2", Reason: "B", UserId: "1"},
		{Id: "3", Reason: "C", UserId: "1"},
	}

	if err := sampleReasons.CreateReasons(); err != nil {
		log.Fatalf("fail: db.Create Reasons, %v\n", err)
	}
}
