package database

import (
	"log"

	"github.com/R0DR160HM/gin-api-rest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	dsn := "root:root@tcp(127.0.0.1:3306)/banco_testes_gim?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
	DB.AutoMigrate(&models.Aluno{})
}
