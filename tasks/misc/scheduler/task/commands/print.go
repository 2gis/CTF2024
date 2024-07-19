package commands

import (
	"fmt"
	"task/entity"
)

type PrintCommand struct {
	Content string
}

func (cmd PrintCommand) GetName() string {
	return "print"
}

func (cmd PrintCommand) Execute(sender entity.Entity) {
	if sender.GetId() != "" {
		fmt.Println(fmt.Sprintf("[%s] > %s", sender.GetId(), cmd.Content))
	}else{
		fmt.Println("Unauthorized!")
	}
}