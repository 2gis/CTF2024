package main

import (
	"fmt"
	"strconv"
	"strings"
)

func check_code(code string) bool {
	parts := strings.Split(code, "-")
	if len(parts) != 5 {
		return false
	}
	if parts[0][:1] != "A" {
		return false
	}
	if len([]rune(parts[0])) != 1 {
		return false
	}
	if parts[1] != "A0A23" {
		return false
	}
	result := 0
	for _, char := range []rune(parts[2]) {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		result += num
	}
	if len([]rune(parts[2])) != 5 {
		return false
	}
	if result != 14 {
		return false
	}
	prev := parts[3][:1]
	for _, char := range []rune(parts[3]) {
		if string(char) != prev {
			return false
		}
	}
	if prev != "B" {
		return false
	}
	if parts[4][3:4] != "X" || parts[4][:1] != "A" || parts[4][1:2] != "F" || parts[4][4:5] != "5" || parts[4][2:3] != "0" {
		return false
	}
	return true
}

func main(){
	var input string
	fmt.Scan(&input)
	if check_code(input) {
		fmt.Println("2GIS.CTF{r3v3r3se_l1c3ns3_w4s_n0t_h4rs_y3ah}")
	}
}
