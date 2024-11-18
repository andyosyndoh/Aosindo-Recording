package internals

import (
	"fmt"
	"strconv"

	"task/models"
)

func AddTask() {
	for {
		fmt.Println("Please Type description Of The task:")
		var Description string
		fmt.Scan(&Description)
		if Description == "" {
			fmt.Println("Description Cannot be Empty, Type Again")
		} else {
			fmt.Println("What Priority do you give Your Task?\n1 High\n2 Medium\n3 Low")
			var Priotity string
			for {
				fmt.Scan(&Priotity)
				num, err := strconv.Atoi(Priotity)
				if err != nil || num < 0 || num > 3 {
					fmt.Println("Please choose a number corresponding to priority:")
				} else {
					Priotity = GetPriority(num)
					id := len(models.Tasks)
					fmt.Printf("Your task id = 00%d with description\n\"%s\"\nand Priority = %v to be added\n", id, Description, Priotity)
					Add(id, Description, num)
					return
				}
			}
		}
	}
}

func GetPriority(num int) string {
	if num == 1 {
		return "High"
	} else if num == 2 {
		return "Madium"
	} else if num == 3 {
		return "Low"
	}
	return "invalid"
}

func  Add(id int, description string, num int) {
	new := models.Task
	new.ID = id
	new.Description = description
	new.Priority = num
	models.Tasks = append(models.Tasks, new)
}
