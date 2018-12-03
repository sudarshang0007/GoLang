package  main

import (
	"fmt"
	_http "net/http"
	_io "io/ioutil"
)

func main(){

	url:= "https://reqres.in/api/users/2";
	var client = _http.Client{};

	// 1. Fetch Simple response from Server
	response, err := client.Get(url);

	if(err != nil){
		fmt.Printf("Message %s",err);
	}else{	
		data, _:=_io.ReadAll(response.Body);
		fmt.Println(response);
		fmt.Println("\n\n body",string(data))

	}



}