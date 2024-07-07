package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	BossBabyRevenge()
}

func BossBabyRevenge() {
	//TODO: split string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	_input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	_input = strings.TrimSpace(_input)

	//TODO: check error input
	if len(_input) == 0 {
		fmt.Println("Error Input")
		return
	}

	if _input[0] == 'R' {
		fmt.Println("Bad boy")
		return
	}

	//TODO: check shots and revenge if short shots += 1 if revenge shots -= 1
	shots := 0
	for _, char := range _input {
		if char == 'S' {
			shots++
		} else if char == 'R' {
			if shots > 0 {
				shots--
			}
		}
	}
	if shots <= 0 {
		fmt.Println("Good boy")
		return
	} else {
		fmt.Println("Bad boy")
		return
	}
}
