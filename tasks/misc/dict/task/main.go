package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var dict = []string{
	"Hello",
	"World",
}

var commands = map[string]func(){
	"menu": menu,
	"word": word,
	"addword": addword,
	"flag": flag,
	"exit": exit,
}

const SECRET_KEY = "ae43cd632d8a314cde6542dec876542"

func sha256sum(bytes []byte) string {
	h := sha256.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func rangeInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func buildKeyboard(commands []string) string {
	keyboard := "Keyboard: \n"
	for _, command := range commands {
		keyboard += fmt.Sprintf("%s_%s\n", command, sha256sum([]byte(command + SECRET_KEY)))
	}
	return keyboard
}

func exit() {
	os.Exit(0)
}

func menu() {
	fmt.Println("Enter command_sign from keyboard\n" + buildKeyboard([]string{
		"menu",
		"word",
		"addword",
		"exit",
	}))
}

func flag() {
	fmt.Println("2GIS.CTF{fl4g_c0mm4nd_s11gn3dddd}")
}

func word() {
	fmt.Println(dict[rangeInt(0, len(dict) - 1)])
}

func addword() {
	fmt.Println("Enter word")
	var cmd string
	fmt.Scan(&cmd)
	fmt.Println(buildKeyboard([]string{"menu", strings.ReplaceAll(cmd, "flag", "")}))
}

func main() {
	fmt.Println("========================\nDict service\n========================\n")
	menu()
	for {
		var cmd string
		fmt.Scan(&cmd)
		found := false
		for command, fn := range commands {
			if strings.HasPrefix(cmd, command) {
				if sha256sum([]byte(command + SECRET_KEY)) == strings.Split(cmd, "_")[1] {
					fn()
				}else{
					fmt.Println("BAD SIGN!")
				}
				found = true
				break
			}
		}
		if !found {
			args := strings.Split(cmd, "_")
			if sha256sum([]byte(args[0]+SECRET_KEY)) == args[1] {
				dict = append(dict, args[0])
				fmt.Println("word added!")
			} else {
				fmt.Println("BAD SIGN!")
			}
		}
	}
}
