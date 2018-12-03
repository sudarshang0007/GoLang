package main

import (
"fmt"
"time"
"strconv"
)

func main() {
    c1 := make(chan string, 1)
    c2 := make(chan string, 1)
    go writer1("writer1",c1)
    go writer2("writer2",c2)
    go reader(c1,c2)
    fmt.Scanln()
}

// Single Reader is used to read the data from two writers
func reader (c1 chan string,c2 chan string){
    for{
        select{
        case msg1 := <- c1:
                fmt.Println(msg1)
        case msg2 := <- c2:
                fmt.Println(msg2)
        }
    }

    /*for msg := range c{
        fmt.Println("In reader ",msg)
        time.Sleep(time.Millisecond * 300)
    }*/
}

func writer1 (str string,c1 chan string){
    for i := 1;i<=5;i++{
        //fmt.Println("In writer ",i)
        c1 <- str 
        time.Sleep(time.Millisecond * 100)
    }
}

func writer2 (str string,c2 chan string){
    for i := 1;i<=5;i++{
        //fmt.Println("In writer ",i)
        c2 <- str 
        time.Sleep(time.Millisecond * 20)
    }
} 