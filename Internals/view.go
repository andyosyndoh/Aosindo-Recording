package internals

import (
	"fmt"
	"task/models"
)


func ViewTasks() {
	if len(models.Tasks) == 0 {
		fmt.Println("You Have no tasks")
		return
	}
	for i :=0; i < len(models.Tasks)-1;i++ {
		if models.Tasks[i+1].Priority < models.Tasks[i].Priority {
			models.Tasks[i], models.Tasks[i+1] = models.Tasks[i+1], models.Tasks[i]
		}
	}
	fmt.Println("Your Tasks Are:")
	count := 0
	for _, v := range models.Tasks {
		count ++
		priority := GetPriority(v.Priority)
		fmt.Printf("Task: %v\n ID = %v\n Description = %v\n Priotity = %v", count, v.ID,v.Description,priority)
	}
}