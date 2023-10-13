package database

import (
	"golang-first-api-rest/models"
)

func CriaPessoa(pessoa *models.Pessoa) error {
	return DB.Create(pessoa).Error
}

func BuscaPessoaPorId(pessoa *models.Pessoa, id string) bool {
	return DB.Where("id = ?", id).First(pessoa).RowsAffected == 0
}

func BuscaPessoasPorTermo(pessoas *[]models.Pessoa, termo string) bool {
	return DB.Where("searchable ILIKE ?", "%"+termo+"%").
		Limit("50").Find(pessoas).RowsAffected == 0
}

func BuscaTotalDePessoasCadastradas(result *int64) {
	DB.Table("pessoas").Count(result)
}
