package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang-first-api-rest/models"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable" //host é o nome do serviço Docker no docker-compose
	DB, err = gorm.Open("postgres", stringDeConexao)
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	DB.LogMode(true)
	//para criar uma tabela no banco de dados com
	//base na struct Pessoa, utilizando gorm
	DB.AutoMigrate(&models.Pessoa{})

	//configuração do pool de conexões
	DB.DB().SetMaxOpenConns(15)
}
