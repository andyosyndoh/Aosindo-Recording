package main

import (
	"fmt"
	"os"
	"strconv"
	internals "task/Internals"
)

func main() {
	var count int
	internals.Clear()
	fmt.Println("Hello Welcome to task manager")
	for {
		fmt.Println("Please choose what you want to do:\n1 Add task\n2 View Task\n3 Delete task\n4 Update task\n5 exit")
		var task string
		fmt.Scan(&task)
		num, err := strconv.Atoi(task)
		if err != nil || num < 1 || num > 5 {
			count++
			fmt.Println("Invalid option, choose Again")
		} else {
			count = 0
			if num == 1 {
				internals.AddTask()
			}
			if num == 2 {
				internals.ViewTasks()
			}
			if num == 3 {
				internals.Delete()
			}
			if num == 4 {
				internals.UpdateTasks()
			}
			if num == 5 {
				fmt.Println("Exiting Program, Bye")
				os.Exit(0)
			}
		}
		if count > 3 {
			fmt.Println("Too many invalid options\n Quiting program...")
			os.Exit(0)
		}
	}
}
