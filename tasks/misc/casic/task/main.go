package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func RangeInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	rand.Seed(time.Now().UnixNano() - 0x3223)
	money := 100
	bet := 0
	fmt.Println("Welcome to Casic!\nCommands:\nflag - buy flag\nplay - play casic\nbet - set bet\nmoney - show your money\nexit - exit\n")
	for {
		fmt.Print("> ")
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "flag" {
			if money >= 5000000000 {
				fmt.Println("2GIS.CTF{c4s1c_m0n3y_pwn33d_by_1nt3g3r_0v3rf10w}")
				money -= 5000000000
			}else{
				fmt.Println("You have no money.")
			}
		} else if cmd == "money" {
			fmt.Printf("Your money: %d\n", money)
		} else if cmd == "bet" {
			fmt.Print("Enter bet\n> ")
			var cmd string
			fmt.Scan(&cmd)
			atoi, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("Incorrect bet.")
				continue
			}
			bet = atoi
			fmt.Println("done.")
		} else if cmd == "play" {
			fmt.Print("Enter number\n> ")
			var cmd string
			fmt.Scan(&cmd)
			num, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("Incorrect number.")
				continue
			}
			if bet <= 0 {
				fmt.Println("Incorrect bet.")
				continue
			}
			winbet := bet * 10
			losebet := bet * 2
			if money < losebet {
				fmt.Println("You have no money.")
				continue
			}
			num2 := RangeInt(0, 1000)
			if num == num2 {
				money += winbet
				fmt.Println("Win!")
			} else {
				money -= losebet
				fmt.Printf("You lose! Number was: %d\n", num2)
			}
		} else if cmd == "exist" {
			os.Exit(0)
		} else {
			fmt.Println("Unknow command.")
		}
	}
}
