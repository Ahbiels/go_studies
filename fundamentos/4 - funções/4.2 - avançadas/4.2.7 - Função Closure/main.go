package main

import (
	"fmt"
	"strconv"
)

func closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func user(name string) func(int) string {
	return func(age int) string {
		age_string := strconv.Itoa(age)
		return "Meu nome é " + name + " e eu tenho " + age_string + " anos"
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	value := closure()

	fmt.Println(value(10))
	fmt.Println(value(10))
	fmt.Println(value(10))

	user1 := user("Angelo")
	fmt.Println(user1(18))
	fmt.Println(user1(19))
	fmt.Println(user1(25))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
