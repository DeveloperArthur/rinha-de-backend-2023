package service

import (
	"golang-first-api-rest/database"
	"golang-first-api-rest/models"
)

func CriaPessoa(pessoa *models.Pessoa) error {
	//https://github.com/DeveloperArthur/rinha-de-backend-2023#indexa%C3%A7%C3%A3o-de-pesquisa-textual
	pessoa.SetSearchable()

	return database.CriaPessoa(pessoa)
}

func BuscaPessoaPorId(pessoa *models.Pessoa, id string) bool {
	return database.BuscaPessoaPorId(pessoa, id)
}

func BuscaPessoasPorTermo(pessoas *[]models.Pessoa, termo string) {
	database.BuscaPessoasPorTermo(pessoas, termo)
}

func BuscaTotalDePessoasCadastradas(result *int64) {
	database.BuscaTotalDePessoasCadastradas(result)
}
