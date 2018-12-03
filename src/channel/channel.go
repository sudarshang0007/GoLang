package main

import ( 
	"fmt"
	"time"
)


func main(){

	c:= make(chan string,1);

	

	for i:=1;i<10;i++ {
		go count("sheep",c);
		fmt.Println("Channel value : ", <- c)
	}

}


func count(str string, c chan string){
	time.Sleep(time.Duration(1000));
	c <- str ;

}
