package commands

import (
	"fmt"
	"task/entity"
	"time"
)

type TimeCommand struct {
}

func (cmd TimeCommand) GetName() string {
	return "time"
}

func (cmd TimeCommand) Execute(sender entity.Entity) {
	if sender.GetId() != "" {
		fmt.Println(fmt.Sprintf("%d", time.Now().Unix()))
	}else{
		fmt.Println("Unauthorized!")
	}
}