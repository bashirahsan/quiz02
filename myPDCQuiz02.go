package quiz02
import (
	"fmt"
)
func MyPackage(Choice int) int {
		Cases := []string{"Pakistan has 1526 cases.","South Korea has 9583 cases.","Farance has 37575 cases.","Go Cororna GO"}
		if Choice == 1 {
			fmt.Println(Cases[0])
			return 1
		}else if Choice == 2 {
			fmt.Println(Cases[1])
			return 2
		}else if Choice == 3 {
			fmt.Println(Cases[2])
			return 3
		}else if Choice == 4 {
			fmt.Println(Cases[3])
			return 4
		}else if Choice == 0 {
			return 0
		}
		fmt.Println("Option not valid.")
		return -1

}