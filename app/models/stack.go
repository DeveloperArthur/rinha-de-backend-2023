package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Stack struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Nome      string    `sql:"index"`
	PessoaID  uuid.UUID `gorm:"type:uuid;column:pessoa_foreign_key;not null;"`
	CreatedAt time.Time
}

// BeforeCreate definirá um UUID em vez de um ID numérico.
func (stack *Stack) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

func GetStackNames(stack *[]Stack) []string {
	var stackNames []string
	for _, item := range *stack {
		stackNames = append(stackNames, item.Nome)
	}
	return stackNames
}

func ConvertStringStackNamesToStackList(stackNames *[]string) []Stack {
	stackList := make([]Stack, len(*stackNames))
	for i, stackName := range *stackNames {
		stackList[i] = Stack{Nome: stackName}
	}
	return stackList
}
