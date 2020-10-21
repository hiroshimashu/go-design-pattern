package main

import (
	"fmt"
)

func main() {
	user1 := NewUserBuilder(0).SetName("mashu_matrix").SetGender("men").SetAge(29).Build()
	user2 := NewUserBuilder(1).SetName("mai").SetGender("women").SetAge(27).Build()

	fmt.Printf("user1: %v\n", user1)
	fmt.Printf("user2: %v", user2)
}
