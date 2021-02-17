package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func printGreeting() {
	fmt.Println("Welcome to Lulu's Virtual Vending Machine!")
	fmt.Println("Choose a snack, press d to add money, or q to exit")
}

// clears the terminal
func clear() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func padString(s string, ln int) string {
	if len(s) >= ln {
		return s
	}

	for  len(s) < ln {
		s += " "
	}

	return s
}





func displaySnacks(snacks *Snacks) {
	pad := snacks.longestName() + 2

	for idx,snack := range snacks.Snacks {
		// make the snacks 3 to a line
		if idx > 0 && idx % 3 == 0 {
			fmt.Print("\n")
		}
		price := fmt.Sprintf("%.2f", snack.Price)
		if snack.Qty <= 0 {
			price = " OUT"
		}
		fmt.Printf("%d: %s $%s\t", idx+1,padString(snack.Name, pad), price)
	}
	fmt.Print("\n")
}

func displaySnack(snack *Snack) {
	fmt.Println(snack.Name)
	fmt.Println(snack.Desc)
}

func inputLine(info string) {
	fmt.Printf("\n%s: ", info)
}


func MainScreen(state *State)  {
	clear()
	printGreeting()
	displaySnacks(state.Snacks)
	line := fmt.Sprintf("Balance: %.2f. Select a snack or press q to quit", state.Balance)
	inputLine(line)
}

func SnackScreen(state *State) {
	clear()
	snack := state.Snacks.Snacks[state.Selection]
	displaySnack(&snack)
	line := fmt.Sprintf("Balance: %.2f", state.Balance)
	if state.Balance < snack.Price {
		 line += " Insufficient funds, press d to add $1, or c to cancel selection: "
	} else {
		line += " Press a to accept, or c to cancel"
	}
	inputLine(line)
}

func DispenseScreen(state *State) {
	clear()
	snack := state.Snacks.Snacks[state.Selection]
	displaySnack(&snack)
	fmt.Println("Thank you for choosing LuLu's Vending Machine")
	fmt.Printf("Your change is %.2f", state.Balance)

	inputLine("Press m for the Snacks menu, or q to quit")
}
