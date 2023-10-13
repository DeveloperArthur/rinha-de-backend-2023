package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Pessoa struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Apelido    string    `gorm:"unique"`
	Nome       string
	Nascimento string
	Stack      pq.StringArray `gorm:"type:text[]"`
	Searchable string         `sql:"index"`
}

func (pessoa *Pessoa) SetSearchable() {
	pessoa.Searchable = pessoa.Nome + pessoa.Apelido
	for _, s := range pessoa.Stack {
		pessoa.Searchable += s
	}
}

// BeforeCreate definirá um UUID em vez de um ID numérico.
func (pessoa *Pessoa) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
