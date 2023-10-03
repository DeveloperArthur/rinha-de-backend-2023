package dto

import (
	uuid "github.com/satori/go.uuid"
	"golang-first-api-rest/models"
)

type PessoaDto struct {
	Id         uuid.UUID `json:"id"`
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento string    `json:"nascimento"`
	Stack      []string  `json:"stack"`
}

func ConvertDtoToModel(pessoaDto *PessoaDto) models.Pessoa {
	return models.Pessoa{
		Apelido:    pessoaDto.Apelido,
		Nome:       pessoaDto.Nome,
		Nascimento: pessoaDto.Nascimento,
		Stack:      pessoaDto.Stack,
	}
}

func ConvertModelToDto(pessoa *models.Pessoa) PessoaDto {
	return PessoaDto{
		Id:         pessoa.ID,
		Apelido:    pessoa.Apelido,
		Nome:       pessoa.Nome,
		Nascimento: pessoa.Nascimento,
		Stack:      pessoa.Stack,
	}
}

func ConvertModelListToDtoList(pessoas *[]models.Pessoa) []PessoaDto {
	pessoasDto := make([]PessoaDto, len(*pessoas))
	for i, pessoa := range *pessoas {
		pessoasDto[i] = ConvertModelToDto(&pessoa)
	}
	return pessoasDto
}
