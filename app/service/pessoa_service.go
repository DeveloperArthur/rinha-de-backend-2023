package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang-first-api-rest/caching"
	"golang-first-api-rest/database"
	"golang-first-api-rest/models"
	"golang-first-api-rest/queue"
	"golang-first-api-rest/util"
)

func CriaPessoa(pessoa *models.Pessoa) {
	//Devemos gerar o Id pois é obrigatório que a API retorne o Id
	//Se não gerarmos o Id nesse momento, a API irá retornar Id null
	//Pois pessoa NÃO será persistida agora, será enfileirada e salva no cache
	pessoa.GenerateId()

	//https://github.com/DeveloperArthur/rinha-de-backend-2023#indexa%C3%A7%C3%A3o-de-pesquisa-textual
	pessoa.SetSearchable()

	queue.Sender(pessoa)
	caching.Set(pessoa, pessoa.ID.String())
}

func BuscaPessoaPorId(pessoa *models.Pessoa, id string) bool {
	var recordNotFound bool
	fmt.Println("LOG: Buscando do cache")
	pessoaEmJson, err := caching.Get(id)
	if err == redis.Nil {
		fmt.Println("LOG: Não está no cache, buscando no banco...")
		recordNotFound = database.BuscaPessoaPorId(pessoa, id)
		if recordNotFound == false {
			caching.Set(pessoa, id)
			fmt.Println("LOG: Salvo no cache com sucesso")
		}
		return recordNotFound
	} else if err != nil {
		panic(err)
	} else {
		//record found, está no cache!
		recordNotFound = false
		util.Deserialize(pessoaEmJson, pessoa)
		return recordNotFound
	}
}

func BuscaPessoasPorTermo(pessoas *[]models.Pessoa, termo string) {
	fmt.Println("LOG: Buscando do cache")
	pessoaEmJson, err := caching.Get(termo)
	if err == redis.Nil {
		fmt.Println("LOG: Não está no cache, buscando no banco...")
		recordNotFound := database.BuscaPessoasPorTermo(pessoas, termo)
		if recordNotFound == false {
			caching.SetList(pessoas, termo)
			fmt.Println("LOG: Salvo no cache com sucesso")
		}
	} else if err != nil {
		panic(err)
	} else {
		//record found, está no cache!
		util.DeserializeList(pessoaEmJson, pessoas)
	}
}

func BuscaTotalDePessoasCadastradas(result *int64) {
	database.BuscaTotalDePessoasCadastradas(result)
}

func PessoaJaExisteNoBanco(pessoa *models.Pessoa) bool {
	return database.PessoaJaExisteNoBanco(pessoa)
}
