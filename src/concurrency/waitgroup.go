package main
import (
	"fmt"
	"strconv"
	"sync"
	 "time"
)


func main(){

	wg:=&sync.WaitGroup{};
	
	wg.Add(1);
	go func(){
		printSeq("first", 1500);
		wg.Done();
	}()

	wg.Add(1);
	go func(){
		printSeq("second", 2000);
		wg.Done();
	}()
	wg.Wait();
}

func printSeq(str string, timer int){
	for i:=1 ;i<10;i++{
		time.Sleep(time.Duration(timer));
		fmt.Println(str + strconv.Itoa(i));
	}
}