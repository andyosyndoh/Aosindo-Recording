package internals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"task/models"
)

func UpdateTasks() {
	scanner := bufio.NewScanner(os.Stdin)
	Clear()
	id := ""
	for {
		fmt.Println("Please enter Task Id to be updtaed:")
		fmt.Scan(&id)
		num, err := strconv.Atoi(id)
		if err != nil || num >= len(models.Tasks) {
			fmt.Println("Invalid Id Please try again")
		} else {
			for _, v := range models.Tasks {
				if v.ID == num {
					for scanner.Scan(){
						fmt.Println("Please Type new description Of The task:")
						Description := scanner.Text()
						if Description == "" {
							fmt.Println("Description Cannot be Empty, Type Again")
						} else {
							fmt.Println("What new Priority do you give Your Task?\n1 High\n2 Medium\n3 Low")
							var Priotity string
							for {
								fmt.Scan(&Priotity)
								num, err := strconv.Atoi(Priotity)
								if err != nil || num < 0 || num > 3 {
									fmt.Println("Please choose a number corresponding to priority:")
								} else {
									Priotity = GetPriority(num)
									id := len(models.Tasks)
									fmt.Printf("Your task id = 00%d with description\n\"%s\"\nand Priority = %v to be updated\ncontinue(yes/no)", id, Description, Priotity)
									consent := ""
									fmt.Scan(&consent)
									if consent == "yes" || "consent" == "Yes" {
										v.Description = Description
										v.Priority = num
										Clear()
										fmt.Println("Updated succesfully")
										return
									} else {
										Clear()
										fmt.Println("Updated aborted")
										return
									}
								}
							}
						}
					}
				}
			}
			fmt.Printf("Id : %v not found", id)
		}
	}
}
