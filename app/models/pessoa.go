package models

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Pessoa struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Apelido    string
	Nome       string
	Nascimento string
	Stack      pq.StringArray `gorm:"type:text[]"`
	Searchable string         `sql:"index"`
}

func (pessoa *Pessoa) GenerateId() {
	pessoa.ID = uuid.NewV4()
}

func (pessoa *Pessoa) SetSearchable() {
	pessoa.Searchable = pessoa.Nome + pessoa.Apelido
	for _, s := range pessoa.Stack {
		pessoa.Searchable += s
	}
}
