package commands

import (
	"fmt"
	"task/entity"
)

const FLAG = "2GIS.CTF{scheduler_privilege_escalation_executed_to_operator_permissions}"

type FlagCommand struct {}

func (cmd FlagCommand) GetName() string {
	return "flag"
}

func (cmd FlagCommand) Execute(sender entity.Entity) {
	if sender.GetId() != "" {
		switch sender.(type) {
		case *entity.User:
			if sender.GetAdminLevel() == 2 {
				fmt.Println(FLAG)
				return
			}else{
				fmt.Println("You have no permission to read flag")
				return
			}
		}
		fmt.Println("You're not user!")
	}else{
		fmt.Println("Unauthorized")
	}
}