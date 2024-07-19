package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"task/commands"
	"task/entity"
)

func Md5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func RangeInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

const SECRET_KEY = "bcd4a36c914450d1bdec98c0ccf366fdfaa762b59d1d570bef518142758ee585a962587e82b4470f1fa8fe51e228f4a442221fc76c04ff987ca9de4f551e18c7"

func main() {
	fmt.Println("Type \"register\" to create your account")
	user := entity.User{}
	for {
		var cmd string
		fmt.Scan(&cmd)
		if strings.HasPrefix(cmd, "register") {
			if user.GetId() == "" {
				user.Id = Md5([]byte(strconv.Itoa(RangeInt(-1 << 31 + 2, 1<<31 - 2)) + SECRET_KEY))
				user.AdminLevel = 1
				entity.Entities[user.Id] = &user
				fmt.Println("Welcome! You can execute commands!\n* time\n* print_{message}\n* schedule_{time}_{command}\n* flag")
			}else{
				fmt.Println("Already registered")
			}
		}else if strings.HasPrefix(cmd, "print_") {
			commands.CommandExecutor{
				Command: commands.PrintCommand{Content: strings.Replace(cmd, "print_", "", 1)},
				ExecutorEntity: &user,
			}.Execute()
		}else if strings.HasPrefix(cmd, "flag") {
			commands.CommandExecutor{
				Command:        commands.FlagCommand{},
				ExecutorEntity: &user,
			}.Execute()
		}else if strings.HasPrefix(cmd, "time") {
			commands.CommandExecutor{
				Command: commands.TimeCommand{},
				ExecutorEntity: &user,
			}.Execute()
		}else if strings.HasPrefix(cmd, "op") {
			opuser := strings.Replace(cmd, "op_", "", 1)
			commands.CommandExecutor{
				Command: commands.OpCommand{entity.GetEntityById(opuser)},
				ExecutorEntity: &user,
			}.Execute()
		}else if strings.HasPrefix(cmd, "schedule_") {
			args := strings.Split(cmd, "_")
			unixtime, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err.Error())
			}
			commands.CommandExecutor{
				Command: commands.ScheduleCommand{Time: int64(unixtime), Cmd: strings.Replace(cmd, fmt.Sprintf("%s_%s_", args[0], args[1]), "", 1)},
				ExecutorEntity: &user,
			}.Execute()
		}
	}
}
