package piscine

import "fmt"

func PrintComb() {
	j:=1
	d:=2
	for i := 0; i <= 5; i++ {
		for j = i + 1; j <= 6; j++ {
			for d = j + 1; d <= 7; d++ {

				fmt.Print(i)
				fmt.Print(j)
				fmt.Print(d)
				if i < 5 {
				fmt.Print(", ")
				}else {
				    fmt.Print("\n")
				}

			}
		}
	}
}