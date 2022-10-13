package main

import "fmt"

func main() {
	var myName = "Kalpesh"
	fmt.Println(`My Name is `, myName, myName)

	var name string = "KC"
	fmt.Println("name = ", name)

	userName := "Admin"
	fmt.Println("username = ", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 6
	fmt.Println("part 1 is ", part1, " other is ", other)

	part2, other := 2, 0
	fmt.Println("part 2 is ",part2," other is ",other)

	sum = part1 + part2
	fmt.Println("Sum = ", sum)

	var (
		lesson = "Variables"
		lessonType = "demo"
	)
	fmt.Println("lessonName = ", lesson)
	fmt.Println("lessonType = ", lessonType)

	word1, word2, _ := "Hello", "world", "!"
	fmt.Println(word1,word2)
}
