package main

import (
	"fmt"
	"github.com/lithammer/shortuuid"
	"log"
	"mariadbapp/config"
	"mariadbapp/database/gormdb"
	kafkaservice "mariadbapp/kafka-service"
	"os"
)

type User struct {
	//gorm.Model
	FirstName      string
	LastName       string
	Email          string `gormdb:"unique_index:user_email_index"`
	Password       string
	Token          string
	TokenExpiresAt uint
	ID             string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	log.New(os.Stdout, "", log.Lmicroseconds|log.Lshortfile)
	c := config.ReadEnv()
	if c == nil {
		log.Println("Failed to get env variables from .yml file")
		return
	}
	db, err := gormdb.NewGormConfig(c)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database!")
	}

	if db.HasTable(&User{}) == false {
		db.CreateTable(&User{})
	}

	user := User{
		FirstName: "john",
		LastName:  "doe",
		Email:     "john.doe@emai.com",
		Password:  "insecurepassword",
	}
	user.ID = shortuuid.New()
	//db.Create(&user)
	db.Model(&User{}).Create(&user)

	producer, err := kafkaservice.InitProducer()
	if err != nil {
		return
	}
	kafkaservice.Publish("new entry added successfully", producer)
	var users []User
	db.Find(&users)

	fmt.Println("There are", len(users), "user records in the table.")

}
