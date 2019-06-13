package piscine

import "fmt"

func PrintComb() {
	j:=1
	d:=2
	for i := 0; i <= 7; i++ {
		for j = i + 1; j <= 8; j++ {
			for d = j + 1; d <= 9; d++ {

				fmt.Print(i)
				fmt.Print(j)
				fmt.Print(d)
				if i < 7 {
				fmt.Print(", ")
				}else {
				    fmt.Print("\n")
				}

			}
		}
	}
}