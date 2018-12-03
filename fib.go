package main

import "fmt"

func main(){
	var newNo=0;
	firstNo:=1; secNo:=1; level:=10;

	fmt.Printf("%d %d ",firstNo,secNo);
	
	for i:=0; i<level ; i++ {
		newNo = firstNo + secNo;
		fmt.Printf("%d " ,newNo);
		firstNo = secNo;
		secNo= newNo;
	}
}