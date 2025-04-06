// Creating my first todolist application by using go language
// Basically it focus on eid day
package main

import "fmt"

func main() {
	cloth := "Buy new cloth for eid."
	clean := "Clean & decorate the house."
	food := "Prepare or order special Eid food."
	wishes := "send Eid wishes to family & friends."
	money := "Set aside money or gifts for Eidi."

	todoItems := []string{cloth, clean, food, wishes, money}

	fmt.Println("***Welcome to our Todolist***")
	fmt.Println()
	fmt.Println("list of todos")
	fmt.Println()

	for index, task := range todoItems {

		fmt.Printf("%d. %s \n", index+1, task) //after using is it well print the result ofc

	}

}
