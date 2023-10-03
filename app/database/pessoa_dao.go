package database

import (
	"golang-first-api-rest/models"
)

func CriaPessoa(pessoa *models.Pessoa) {
	DB.Create(pessoa)
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

func PessoaJaExisteNoBanco(pessoa *models.Pessoa) bool {
	result := DB.Where(&models.Pessoa{Apelido: pessoa.Apelido}).First(pessoa)
	if result.RecordNotFound() {
		return false
	}
	return true
}
