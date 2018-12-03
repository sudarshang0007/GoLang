package main

import (
	"fmt"
	"os" 
)

func main() {
	a := 20;
	fmt.Printf("Hello World !! %T",a)
	for i:=1 ;  i < len(os.Args); i++ {
		fmt.Printf("\n %s",os.Args[i]);
	}
}
