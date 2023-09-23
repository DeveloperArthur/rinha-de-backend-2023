package util

import (
	"encoding/json"
	"golang-first-api-rest/models"
)

func Serialize(pessoa *models.Pessoa) string {
	pessoaEmJson, err := json.Marshal(pessoa)
	if err != nil {
		panic(err)
	}
	return string(pessoaEmJson)
}

func SerializeList(pessoa *[]models.Pessoa) string {
	pessoaEmJson, err := json.Marshal(pessoa)
	if err != nil {
		panic(err)
	}
	return string(pessoaEmJson)
}

func Deserialize(pessoaEmJson string, pessoa *models.Pessoa) {
	err := json.Unmarshal([]byte(pessoaEmJson), &pessoa)
	if err != nil {
		panic(err)
	}
}

func DeserializeList(pessoaEmJson string, pessoas *[]models.Pessoa) {
	err := json.Unmarshal([]byte(pessoaEmJson), &pessoas)
	if err != nil {
		panic(err)
	}
}
