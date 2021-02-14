package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


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

func validateSelection(choice string, snacks *Snacks) bool {
	num, err := strconv.Atoi(choice)

	if err != nil {
		fmt.Println("NOT A NUMBER!!!!!")
		return false
	}

	if num >= len(snacks.Snacks) || num <= 0 {
		return false
	}

	if snacks.Snacks[num-1].Qty <= 0 {
		return false
	}

	return true

}

func printGreeting() {
	fmt.Println("Welcome to Lulu's Virtual Vending Machine!")
	fmt.Println("Choose a snack, press d to add money,  or press Q to exit")
}

func getUserInput(reader *bufio.Reader) string {
	fmt.Print("Make a selection: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.Replace(choice, "\n", "", -1)
	return choice
}

func main() {
	path := "snacks.json"
	snacks,err := GetSnacks(path)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	printGreeting()
	for {

		displaySnacks(snacks)

		choice := getUserInput(reader)

		if choice == "q" || choice == "Q" {
			fmt.Println("Powering down, have a nice day")
			break
		}

		valid := validateSelection(choice, snacks)

		if valid {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid selection, chose an available snack or press q to exit")
			continue
		}
	}



}


