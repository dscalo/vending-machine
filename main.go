package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)


func getSelection(choice string, snacks *Snacks) int {
	num, err := strconv.Atoi(choice)

	if err != nil {
		return -1
	}

	if num >= len(snacks.Snacks) || num <= 0 {
		return -1
	}

	if snacks.Snacks[num-1].Qty <= 0 {
		return -1
	}

	return num - 1

}

func getUserInput(reader *bufio.Reader) string {
	choice, _ := reader.ReadString('\n')
	choice = strings.Replace(choice, "\n", "", -1)
	return strings.ToLower(choice)
}


func processSelection(state *State, action string)  {
	switch state.Screen {
		case "MAIN":
			if action == "q" {
				state.Screen = "QUIT"
			}
			state.Selection = getSelection(action, state.Snacks)
			if state.Selection == -1 {
				return
			}
			state.Screen = "SNACK"

		case "SNACK":
			switch action {
				case "d":
					state.Balance += 1.00
					if state.Balance >= state.Snacks.Snacks[state.Selection].Price {
						state.Screen = "CONFORMATION"
					}
				case "c":
					state.Screen = "MAIN"
					state.Selection = -1
				case "q":
					state.Screen = "QUIT"
			}

		case "CONFORMATION":
			switch action {
				case "c":
					state.Screen = "MAIN"
					state.Selection = -1
				case "a":
					state.Screen = "DISPENSE"
					state.Balance -= state.Snacks.Snacks[state.Selection].Price
				case "q":
					state.Screen = "QUIT"
			}
		case "DISPENSE":
			switch action {
				case "m":
					state.Screen = "MAIN"
					state.Selection = -1
					state.Balance = 0
				case "q":
					state.Screen = "QUIT"
			}
	default:
		panic("Unknown screen "+ state.Screen)
		}

}

func main() {
	path := "snacks.json"
	snacks,err := GetSnacks(path)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	state := newState("MAIN", 0.00, snacks)

	for state.Screen != "QUIT" {
		choice := ""
		switch state.Screen {
		case "MAIN" :
			MainScreen(state)
			choice = getUserInput(reader)
		case "SNACK":
			fallthrough
		case "CONFORMATION":
			SnackScreen(state)
			choice = getUserInput(reader)
		case "DISPENSE":
			DispenseScreen(state)
			choice = getUserInput(reader)

		}
		processSelection(state, choice)
	}
}


