package database

import (
	"golang-first-api-rest/models"
)

func CriaPessoa(pessoa *models.Pessoa) {
	DB.Create(pessoa)
}

func BuscaPessoaPorId(pessoa *models.Pessoa, id string) bool {
	return DB.Preload("Stack").
		Where("id = ?", id).First(pessoa).RowsAffected == 0
}

func BuscaPessoasPorTermo(pessoas *[]models.Pessoa, termo string) bool {
	/*
		query:
		SELECT DISTINCT p.id, p.apelido, p.nome, p.nascimento
		FROM pessoas p INNER JOIN stacks s
		ON p.id = s.pessoa_foreign_key
		WHERE p.apelido ILIKE '%termo%'
		OR p.nome ILIKE '%termo%'
		OR s.nome ILIKE '%termo%'
	*/

	return DB.Preload("Stack").Model("pessoas").
		Select("DISTINCT pessoas.id, pessoas.apelido, pessoas.nome, pessoas.nascimento").
		Joins("INNER JOIN stacks ON pessoas.id = stacks.pessoa_foreign_key").
		Where("pessoas.apelido ILIKE ?", "%"+termo+"%").
		Or("pessoas.nome ILIKE ?", "%"+termo+"%").
		Or("stacks.nome ILIKE ?", "%"+termo+"%").
		Limit("50").
		Find(pessoas).
		RowsAffected == 0
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
