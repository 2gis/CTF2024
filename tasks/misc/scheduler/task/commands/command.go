package commands

import (
	"task/entity"
)

type Command interface {
	GetName() string
	Execute(entity.Entity)
}