package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)
	fmt.Print("My random number is ", randomNumber)
	if randomNumber >= 50 {
		fmt.Println(" and is greater than 50")
	} else {
		fmt.Println("and is less than 50")
	}
	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Println("I am a student of the", school)
	} else {
		fmt.Println("I am not a student of the", school)
	}
	beautifulWeather := true
	if beautifulWeather {
		fmt.Println("It's beautiful weather!")
	} else {
		fmt.Println("It's not so nice outside")
	}
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	for _, holbertonFounders := range holbertonFounders {
		fmt.Println(holbertonFounders, "is a founder")
	}
}
