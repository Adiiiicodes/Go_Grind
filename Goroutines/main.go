package main

import (
    "fmt"
  
)

/*  // A function that runs as a goroutine
func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    // Start a goroutine
    go sayHello()
    
    // Start an anonymous goroutine
    go func() {
        fmt.Println("Hello from anonymous goroutine")
    }()
    
    // Wait for goroutines to finish
    time.Sleep(1 * time.Second)
    fmt.Println("Hello from main")
} */

func main() {

for i := 0; i < 1000000; i++ {

	go func(n int) {
        fmt.Println("i: ", n)
    }(i)
	

}


// Wait for goroutines to finish
    var input string
    fmt.Scanln(&input)
} 

  


