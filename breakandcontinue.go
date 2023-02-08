package main

import "fmt"

func main() {
	fmt.Println("---break---")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}
	fmt.Println("----break label---")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i:%d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}

	fmt.Println("----continue---")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}
	fmt.Println("----continue label---")
ne:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i:%d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue ne
		}
	}

	fmt.Printf("------GOTO------")
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a is : %d \n", a)
		a++
	}

}
