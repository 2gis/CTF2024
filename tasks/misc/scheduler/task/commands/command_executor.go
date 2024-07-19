package commands

import "task/entity"

type CommandExecutor struct {
	Command Command
	ExecutorEntity entity.Entity
}

func (executor CommandExecutor) Execute() {
	executor.Command.Execute(executor.ExecutorEntity)
}