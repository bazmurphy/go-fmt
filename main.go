package main

import "fmt"

func main() {
	// Printf
	year := 2024
	fmt.Printf("I was born in %d", year) // no space or newline is added
	// I was born in 2024

	// Println
	fmt.Println("Hello", "World") // spaces are added between arguments
	fmt.Println("Hello World")    // newline is appended at the end
	// Hello World
	// Hello World

	// Sprint
	a := fmt.Sprint("Hello", "World") // no space or newline is added
	b := fmt.Sprint("Hello World")
	fmt.Println(a)
	fmt.Println(b)
	// HelloWorld
	// Hello World

	// Sprintln
	c := fmt.Sprintln("Hello", "World")
	fmt.Println(c)
	// Hello World

	// Sprintf
	name := "Baz"
	s := fmt.Sprintf("My name is %s\n", name)
	fmt.Print(s)
	// My name is Baz

	// Scan
	var name2 string
	fmt.Println("What's your name?")
	fmt.Scan(&name2)
	fmt.Println("Nice to meet you", name2)
	// What's your name?
	// Baz
	// Nice to meet you Baz

	// Scanf
	var name3 string
	fmt.Printf("What's your name? : ")
	fmt.Scanf("%s", &name3) // user input from standard input
	fmt.Printf("Your name is %s\n", name3)
	// What's your name? : Baz
	// Your name is Baz

	// Scanln
	var name4 string
	var age int
	fmt.Printf("What's your name and age? : ")
	fmt.Scanln(&name4, &age) // they must be entered with a space between
	fmt.Printf("Your name is %s and you are %d years old\n", name, age)

	// What's your name and age? : Baz 99
	// Your name is Baz and you are 99 years old
}
