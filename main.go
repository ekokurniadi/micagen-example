package main

import (
	"log"
	"tesss/entity"

	"github.com/ekokurniadi/micagen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/mica_generator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	gen := micagen.Micagen{}
	gen.GenerateAll(db, &entity.Customer{})

}
