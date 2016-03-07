package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	fmt.Println("Hello, we are Holberton School")
	fmt.Println("the date is", time.Now().Format(time.RFC850))
	fmt.Println("the year is", t.Year())
	fmt.Println("the month is", t.Month())
	fmt.Println("the day is", t.Day())
}
