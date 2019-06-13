package piscine

import "fmt"
import "strconv"

func PrintComb2()  {
	for a:=0; a<10; a++ {
		for b:=0; b<10; b++ {
			for c:=0; c<10; c++ {
				for d:=0; d<10; d++ {
					nb1 := strconv.Itoa(a) + strconv.Itoa(b)
					nb2 := strconv.Itoa(c) + strconv.Itoa(d)
					if(nb1 < nb2){
						fmt.Print(nb1 + " " + nb2)
						if nb1 < "98" {
							fmt.Print(", ")
						}else {
							fmt.Print("\n")
						}
					}
				}
			}
		}
	}
}