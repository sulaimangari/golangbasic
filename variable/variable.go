package main

import "fmt"

func main() {
	var name string

	name = "Marcus"
	fmt.Println(name)

	name = "Polax"
	fmt.Println(name)

	var age = 15
	fmt.Println(age)

	var height int32 = 180
	fmt.Println(height)

	homeTown := "texas"
	fmt.Println(homeTown)

	var (
		firstName = "Marcus"
		lastName  = "Polax"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
