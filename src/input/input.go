package main

import (
	"fmt"
	"log"
)

func main(){

	var ans int;

	fmt.Printf("Enter Number :\n");
	aa, err:=fmt.Scanln( &ans );

	if(err != nil){
		log.Fatal(aa)
	}else{
		fmt.Printf("Answer is %d", ans*ans);
	}
	


}