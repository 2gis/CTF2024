package commands

import (
	"fmt"
	"task/entity"
)

type OpCommand struct {
	Entity entity.Entity
}

func (cmd OpCommand) GetName() string {
	return "op"
}

func (cmd OpCommand) Execute(sender entity.Entity) {
	if sender.GetId() != "" {
		if sender.GetAdminLevel() == 2 {
			cmd.Entity.SetAdminLevel(2)
			fmt.Println("Op!")
		}else{
			fmt.Println("You have no permission to do this!")
		}
	}else{
		fmt.Println("Unauthorized!")
	}
}