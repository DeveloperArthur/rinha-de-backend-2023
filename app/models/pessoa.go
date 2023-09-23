package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Pessoa struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Apelido    string    `sql:"index"`
	Nome       string    `sql:"index"`
	Nascimento string
	Stack      []Stack
	CreatedAt  time.Time
}

func (pessoa *Pessoa) GenerateId() {
	pessoa.ID = uuid.NewV4()
}
