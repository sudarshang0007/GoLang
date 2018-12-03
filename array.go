package main

import "fmt"

func main(){
	var arr[20] int;
	arr[0]=1; arr[1]=1;
	fmt.Printf("%d ", arr[0]);
	fmt.Printf("%d ", arr[1]);
	
	for i:=2; i<len(arr) ; i++ {
		arr[i] = arr[i-1] + arr[i-2];
		fmt.Printf("%d ", arr[i]);
	}
}