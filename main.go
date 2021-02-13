package main

import "fmt"


func main() {
	fmt.Println("Welcome to Lulu's Virtual Vending Machine!")

	path := "snacks.json"
	snacks,err := GetSnacks(path)

	if err != nil {
		panic(err)
	}

	for _,snack := range snacks.Snacks {
		fmt.Println(snack)
	}


}


