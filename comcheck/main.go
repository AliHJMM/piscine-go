package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args[1:]
	for _,chr := range args{
		if chr == "01" || chr == "galaxy" || chr == "galaxy 01"{
			fmt.Print("Alert!!")
		} else{
			return
		} 
	}
	fmt.Print("\n")
}