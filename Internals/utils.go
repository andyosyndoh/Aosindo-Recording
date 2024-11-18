package internals

import "fmt"

func Clear() {
	fmt.Print("\033c")
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