package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(frame string, lname string) (s string) {
	s = fmt.Sprintf(frame, lname)
	return
}
