package main

import ( 
	"fmt"
	"time"
)

func main(){

	// 1.Create Channel
	// 2. Read Input from Console
	// 3. If Input is Integer use First Reader
	// 4. If Input is string  use second Reader



}

// Function Manual Write
func manualWrite(){
	c:= make(chan string,1);
	for i:=1;i<10;i++ {
		go count("sheep",c);
		fmt.Println("Channel value : ", <- c)
	}
}
// function to write on channel
func count(str string, c chan string){
	time.Sleep(time.Duration(1000));
	c <- str ;
}


func reader(){

}

func writer(){

}

func writer1(){

}
