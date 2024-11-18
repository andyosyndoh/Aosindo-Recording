package internals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"task/models"
)

func AddTask() {
	scanner := bufio.NewScanner(os.Stdin)
	Clear()
	fmt.Println("Please Type description Of The task:")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			fmt.Println("Description Cannot be Empty, Type Again")
		} else {
			fmt.Println("What Priority do you give Your Task?(choose number)\n1 High\n2 Medium\n3 Low")
			var Priotity string
			for {
				fmt.Scan(&Priotity)
				num, err := strconv.Atoi(Priotity)
				if err != nil || num < 0 || num > 3 {
					fmt.Println("Please choose a number corresponding to priority:")
				} else {
					Priotity = GetPriority(num)
					id := models.Counter
					fmt.Printf("Your task id = %d with description\n\"%s\"\nand Priority = %v to be added\ncontinue?(yes/no)", id, input, Priotity)
					consent := ""
					fmt.Scan(&consent)
					if consent == "yes" || consent == "Yes" {
						Add(id, input, num)
						models.Counter++
						Clear()
						fmt.Println("Task Added Successfully")
						return
					}

				}
			}
		}
	}
}

func Add(id int, description string, num int) {
	new := models.Task
	new.ID = id
	new.Description = description
	new.Priority = num
	models.Tasks = append(models.Tasks, new)
}
