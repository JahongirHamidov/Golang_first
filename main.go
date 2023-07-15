// package main

// import "fmt"

// func greeting(name string) string {
// 	return "Hello " + name
// }

// func getSum(num1 int, num2 int) int {
// 	return num1 + num2
// }

// func main() {
// 	fmt.Println(greeting("Jakhongir"))
// 	fmt.Println(getSum(3, 4))
// }

package main

import (
	"fmt"
)

func main() {
	//arrays
	// var fruitArr [2]string
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"
	// fmt.Println(fruitArr)

	// fruitArr := [3]string{"apple", "orange", "watermellon"}
	// fmt.Println(fruitArr)

	// slices
	//fruitSlice := []string{"apple", "orange", "grape"}
	//fmt.Println(len(fruitSlice))

	// conditionals

	var x = 10
	var y = 10

	if x <= y {
		fmt.Printf("%d is less or equal than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)

	}

	color := "grey"

	if color == "red" {
		fmt.Println("color is red")
	} else {
		fmt.Println("color is not red")
	}

	//SwitchCase

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	case "green":
		fmt.Println("Color is green")
	default:
		fmt.Println("color is not green, red or blue")
	}

	// Loops

	//Long method
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	//Short method
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number %d  ", i)
	}

	//FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Print(i, " ")
		}
	}

	// Maps

	//Define maps
	contacts := make(map[string]string)

	//assign key values
	contacts["Bob"] = "bob@gmail.com"
	contacts["James"] = "james@gmail.com"
	contacts["Mike"] = "mike@gmail.com"

	fmt.Println(contacts)
	fmt.Println(len(contacts))
	fmt.Println(contacts["Bob"])

	//delete from map
	delete(contacts, "Bob")
	fmt.Println(contacts)

	//Declare map and add key value
	newEmails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(newEmails)
	newEmails["Mike"] = "mike@gmail.com"
	fmt.Println(newEmails)

	// Range
	ids := []int{1, 2, 3, 4, 5, 66}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum= ", sum)

	//Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("\n%s: %s", k, v)
	}

	for k := range emails {
		fmt.Println("\nName: " + k)
	}

	// Pointers
	a := 5
	b := &a

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%T\n", b)

	//Use * to read val from address
	fmt.Println(*b)
	//Change val with pointer
	*b = 10
	fmt.Println(a)

}
