package main

import (
	"fmt"
	"math"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	birth, _ := time.Parse("January 2, 2006", u.DOB)
	age := int(math.Floor(time.Since(birth).Hours() / 8760))
	fmt.Println("Hello", u.Name)
	fmt.Println(u.Name, "who was born in", u.City, "would be", age, "years old today.")
}
