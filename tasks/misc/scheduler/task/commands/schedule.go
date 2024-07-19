package commands

import (
	"strings"
	"task/entity"
	"time"
)

type ScheduleCommand struct {
	Time int64
	Cmd string
}

func (cmd ScheduleCommand) GetName() string {
	return "schedule"
}

func (cmd ScheduleCommand) Execute(sender entity.Entity) {
	var command Command
	if strings.HasPrefix(cmd.Cmd, "print_") {
		command = PrintCommand{Content: strings.Replace(cmd.Cmd, "print_", "", 1)}
	}else if strings.HasPrefix(cmd.Cmd, "flag") {
		command = FlagCommand{}
	}else if strings.HasPrefix(cmd.Cmd, "op_") {
		opuser := strings.Replace(cmd.Cmd, "op_", "", 1)
		command = OpCommand{entity.GetEntityById(opuser)}
	}
	go func() {
		for {
			if time.Now().Unix() >= cmd.Time {
				CommandExecutor{
					Command: command,
					ExecutorEntity: entity.GetEntityById("server"),
				}.Execute()
				return
			}else{
				time.Sleep(time.Second)
			}
		}
	}()
}