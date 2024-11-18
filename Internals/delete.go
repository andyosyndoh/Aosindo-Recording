package internals

import (
	"fmt"
	"strconv"

	"task/models"
)

func Delete() {
	id := ""
	for {
		fmt.Println("Please enter Task Id to be deleted:")
		fmt.Scan(&id)
		num, err := strconv.Atoi(id)
		if err != nil || num >= len(models.Tasks) {
			fmt.Println("Invalid Id Please try again")
		} else {
			for i, v := range models.Tasks {
				if v.ID == num {
					priority := GetPriority(v.Priority)
					fmt.Printf("The following Data is to be deleted:\n Id = %v\n Desciption = %v\n Priority = %v\n continue?(yes/no)\n", v.ID, v.Description, priority)
					consent := ""
					fmt.Scan(&consent)
					if consent == "yes" || consent == "Yes" {
						models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
						fmt.Println("Delete Complete")
						return
					} else {
						fmt.Println("Delete Aborted")
						return
					}
				}
			}
			fmt.Printf("Id : %v not found", id)
		}
	}
}
