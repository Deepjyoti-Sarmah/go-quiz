package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to quiz game")

	fmt.Printf("Enter your name: \n")
	var name string
	fmt.Scan(&name)

	fmt.Printf("hello , %v , welcome to the game! \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	num := 2

	fmt.Printf("What is better the RTX 3080 or RTX 3090? ")
	var ans string
	var ans2 string
	fmt.Scan(&ans, &ans2)
	fmt.Println(ans, ans2)

	if ans+" "+ans2 == "RTX 3090" || ans+" "+ans2 == "rtx 3090" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Println("How many cores does the Ryzen 9 3900x have ? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("You scored %v out of %v", score, num)

	percentage := (float64(score) / float64(num)) * 100

	fmt.Printf("You scored: %v%%", percentage)

}
